package userService

import (
	"backend/middlewares"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"time"
)

type httpRequestKey struct{}

var refreshTokenName = "refresh_token"
var accessExpire = 15 * time.Minute
var refreshExpire = 7 * 24 * time.Hour

// トークンを生成
func generateTokens(c context.Context, id primitive.ObjectID) (*string, error) {
	// アクセストークン
	accessClaims := middlewares.Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpire)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	// リフレッシュトークン
	err = resetRefreshToken(c, id)
	if err != nil {
		return nil, err
	}

	return &accessString, nil
}

// リフレッシュトークンを使用してアクセストークンを再生成
func refreshHandler(c context.Context) (*string, error) {
	// context.Contextからhttp.Requestを取得
	req, ok := c.Value(httpRequestKey{}).(*http.Request)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve *http.Request from context")
	}

	// クッキーを取得
	cookie, err := req.Cookie(refreshTokenName)
	if err != nil {
		return nil, errors.New("refresh token not found")
	}

	// リフレッシュトークンの検証
	token, err := jwt.ParseWithClaims(cookie.Value, &middlewares.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_REFRESH_KEY"), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(*middlewares.Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	// 新しいトークンを発行
	accessToken, err := generateTokens(c, claims.Id)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// リフレッシュトークンを再発行する
func resetRefreshToken(c context.Context, id primitive.ObjectID) error {
	// リフレッシュトークン
	refreshClaims := middlewares.Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpire * time.Second)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(os.Getenv("JWT_REFRESH_KEY"))
	if err != nil {
		return errors.New("invalid or expired refresh token")
	}
	// クッキーの設定
	writer := c.Value("httpResponseWriter").(http.ResponseWriter)
	http.SetCookie(writer, &http.Cookie{
		Name:     refreshTokenName,
		Value:    refreshString,
		Expires:  time.Now().Add(refreshExpire),
		Path:     "/",
		Domain:   "",
		Secure:   false, // TODO: trueにする
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return nil
}
