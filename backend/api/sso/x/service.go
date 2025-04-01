package x

import (
	"backend/db/repository/userRepository"
	"backend/di/handlers"
	"backend/middlewares/customError"
	"backend/service/authService"
	"backend/util/helper"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

func getUserData(c *gin.Context) (*TwitterUser, *customError.Error) {
	//stateのチェック TODO: Redisで実装予定
	code := c.Query("code")
	if code == "" {
		return nil, errMissCode()
	}

	config := NewOAuthConfig()
	// GinのコンテキストからGoのcontext.Contextを取得
	ctx := c.Request.Context()
	token, err := config.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", "test_code_verifier"))
	if err != nil {
		return nil, errMissExchangeToken(err)
	}
	client := config.Client(c, token)
	resp, err := client.Get("https://api.twitter.com/2/users/me?user.fields=profile_image_url")
	if err != nil {
		return nil, errGetUserInfo(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errReadResponseBody(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errGetUserInfoBadStatus(body)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errUnMarshalResponseBody(err)
	}

	// ネストされた `data` の中身だけ取得
	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return nil, errGetData(result)
	}

	// 必要な情報を取り出す
	userInfo := &TwitterUser{
		ID:    data["id"].(string),
		Name:  data["name"].(string),
		Image: data["profile_image_url"].(string),
	}

	return userInfo, nil
}

func createNewUser(c *gin.Context, h *handlers.Handlers, xUser *TwitterUser) (*userRepository.Model, *customError.Error) {
	//画像データを取得する
	img, err := helper.FetchImageFromURL(xUser.Image)
	if err != nil {
		return nil, err
	}
	base64, err := helper.ImageToBase64(img, helper.GenerateBase64Option(100, 100))
	if err != nil {
		return nil, err
	}
	//twitterの情報からユーザーを作成する
	user := &userRepository.Model{
		ID:          primitive.NewObjectID(),
		Name:        xUser.Name,
		Email:       nil,
		Password:    []byte(helper.RandomStr(8)), // 暫定で入れておく(が、ハッシュ化してないので無意味な値)
		TwitterId:   &xUser.ID,
		ImageBase64: base64,
	}
	newUser, err := h.UserHandler.UserRepo.Register(c, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func createUserAndLogin(c *gin.Context, h *handlers.Handlers, xUser *TwitterUser) (*authService.UserWithToken, *customError.Error) {
	newUser, err := createNewUser(c, h, xUser)
	if err != nil {
		return nil, err
	}
	res, err := authService.LoginByUser(newUser, c.Writer, *h.TokenConfig)
	if err != nil {
		return nil, err
	}
	return res, nil
}
