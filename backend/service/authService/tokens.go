package authService

import (
	"backend/middlewares/auth"
	"backend/service/authService/tokenConfig"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var refreshTokenName = "refresh_token"

func parseRefreshToken(req *http.Request, tokenConfig tokenConfig.TokenConfig) (*auth.Claims, error) {
	// クッキーを取得
	cookie, err := req.Cookie(refreshTokenName)
	if err != nil {
		return nil, errors.New("refresh token not found")
	}

	// リフレッシュトークンの検証
	token, err := jwt.ParseWithClaims(cookie.Value, &auth.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return tokenConfig.RefreshSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

// リフレッシュトークンを使用してアクセストークンを再生成
func refreshHandler(req *http.Request, writer http.ResponseWriter, tokenConfig tokenConfig.TokenConfig) (*string, error) {
	claims, err := parseRefreshToken(req, tokenConfig)
	if err != nil {
		return nil, errors.New("invalid claims")
	}

	// 新しいトークンを発行・リフレッシュトークンの再生成
	accessToken, err := GenerateTokens(writer, claims.Id, tokenConfig)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// リフレッシュトークンを再発行する
func resetRefreshToken(writer http.ResponseWriter, id primitive.ObjectID, tokenConfig tokenConfig.TokenConfig) error {
	// リフレッシュトークン
	refreshClaims := auth.Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenConfig.RefreshExpire)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(tokenConfig.RefreshSecretKey)
	if err != nil {
		return errors.New("invalid or expired refresh token")
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     refreshTokenName,
		Value:    refreshString,
		Expires:  time.Now().Add(tokenConfig.RefreshExpire),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}
