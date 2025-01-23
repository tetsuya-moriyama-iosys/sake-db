package x

import (
	"backend/di/handlers"
	"backend/service/userService"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
)

// oauthStateString用にランダムな文字列を生成
func generateStateString() string {
	//TODO:JWTトークンにする案がある
	bytes := make([]byte, 16) // 16バイトのランダムなデータ
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate random bytes: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

// GenerateAuthURL 認証用のURLを生成
func GenerateAuthURL() (*string, error) {
	state := generateStateString()
	config := NewOAuthConfig()
	url, err := config.GenerateAuthCodeURL(state)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func Login(h *handlers.Handlers, c *gin.Context) (*userService.UserWithToken, error) {
	//未ログインパターン(既ログインでも後勝ちでJWT発行する)
	//①新規ユーザー・未ログイン
	//②未ログインで、twitter連携済の既存ユーザー

	//ログイン済パターン→この導線は一旦ないものとして考える TODO: 別APIで実装してみる？
	//①ログイン中のアカウントにtwitterのない既存ユーザー
	//②ログイン中のアカウントに該当twitter連携済既存ユーザー

	//異常系
	//①別アカウントのtwitter連携済ユーザー→連携解除してから再度やってもらう。
	//②ログイン中のアカウント以外に該当twitter連携済のユーザー→連携解除してから再度やってもらう。

	//APIからユーザーデータを取得する
	xUser, err := getUserData(c)
	if err != nil {
		return nil, err
	}

	// twitterIDが該当するユーザーがいるか確認、いればそのユーザーでログイン
	user, err := h.UserHandler.UserRepo.GetByTwitterId(c.Request.Context(), xUser.ID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		// 存在すれば、ログイン
		token, err := userService.GetJWTToken(user)
		if err != nil {
			return nil, err
		}
		return &userService.UserWithToken{
			User:  user,
			Token: token,
		}, nil
	}
	// 存在しなければ、ユーザーを作成してログイン
	newUser, err := createUserAndLogin(c, h.UserHandler.UserRepo, xUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
