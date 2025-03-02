package x

import (
	"backend/db/repository/userRepository"
	"backend/di/handlers"
	"backend/service/authService"
	"backend/util/helper"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

func getUserData(c *gin.Context) (*TwitterUser, error) {
	//stateのチェック TODO: Redisで実装予定
	code := c.Query("code")
	if code == "" {
		return nil, errors.New("missing code")
	}

	config := NewOAuthConfig()
	// GinのコンテキストからGoのcontext.Contextを取得
	ctx := c.Request.Context()
	token, err := config.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", "test_code_verifier")) //TODO:
	if err != nil {
		return nil, errors.New("failed to exchange token")
	}
	client := config.Client(c, token)
	resp, err := client.Get("https://api.twitter.com/2/users/me?user.fields=profile_image_url")
	if err != nil {
		return nil, errors.New("failed to fetch user info")
	}
	defer resp.Body.Close()

	var userInfo TwitterUser
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, errors.New("failed to parse user info")
	}
	return &userInfo, nil
}

func createNewUser(c *gin.Context, h *handlers.Handlers, xUser *TwitterUser) (*userRepository.Model, error) {
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

func createUserAndLogin(c *gin.Context, h *handlers.Handlers, xUser *TwitterUser) (*authService.UserWithToken, error) {
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
