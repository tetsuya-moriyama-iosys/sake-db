package authService

import (
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"backend/middlewares/auth"
	"backend/service/authService/tokenConfig"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func Login(ctx context.Context, writer *http.ResponseWriter, input graphModel.LoginInput, r *userRepository.UsersRepository, tokenConfig tokenConfig.TokenConfig) (*UserWithToken, error) {
	// ユーザーインスタンスを取得
	user, err := getUserByInput(ctx, input, r)
	if err != nil {
		return nil, err
	}

	// JWTトークン生成
	accessToken, err := GenerateTokens(writer, user.ID, tokenConfig)
	if err != nil {
		return nil, err
	}
	return generateUserWithToken(user, *accessToken), nil
}

// RefreshTokens アクセストークンが切れたため、リフレッシュトークンを使いトークンを再生成
func RefreshTokens(req *http.Request, writer *http.ResponseWriter, tokenConfig tokenConfig.TokenConfig) (*string, error) {
	return refreshHandler(req, writer, tokenConfig)
}

// LoginWithRefreshToken リフレッシュトークンを用いてログインする
func LoginWithRefreshToken(ctx context.Context, req *http.Request, writer *http.ResponseWriter, tokenConfig tokenConfig.TokenConfig, r *userRepository.UsersRepository) (*UserWithToken, error) {
	claims, err := parseRefreshToken(req, tokenConfig)
	if err != nil {
		return nil, err
	}

	// ユーザーインスタンスを取得
	return loginById(ctx, claims.Id, writer, tokenConfig, r)
}

// GenerateTokens トークンを生成
func GenerateTokens(writer *http.ResponseWriter, id primitive.ObjectID, tokenConfig tokenConfig.TokenConfig) (*string, error) {
	// アクセストークン
	accessClaims := auth.Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenConfig.AccessExpire)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(tokenConfig.AccessSecretKey)
	if err != nil {
		return nil, err
	}

	// リフレッシュトークン
	err = resetRefreshToken(*writer, id, tokenConfig)
	if err != nil {
		return nil, err
	}

	return &accessString, nil
}

func DeleteRefreshToken(writer http.ResponseWriter) error {
	//クッキーを消去
	http.SetCookie(writer, &http.Cookie{
		Name:     refreshTokenName,
		Value:    "",
		Expires:  time.Unix(0, 0), // 過去の時刻に設定
		MaxAge:   -1,              // 即座に削除
		HttpOnly: true,
	})
	return nil
}
