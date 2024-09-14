package graph

import (
	"backend/middlewares"
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// カスタムディレクティブの実装
func authDirective(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	// ヘッダーからトークンを取得
	tokenString, err := extractTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// トークンをパースして検証
	claims := &middlewares.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return middlewares.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// ユーザーIDをcontextに保存
	ctx = context.WithValue(ctx, middlewares.UserContextKey, claims.Id)

	// 認証に成功した場合、次のリゾルバを実行
	return next(ctx)
}

// JWTトークンを読み込むための関数
func extractTokenFromHeader(ctx context.Context) (string, error) {
	req := ctx.Value("http.Request").(*http.Request)
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	// "Bearer "を除去してトークンを取得
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return "", errors.New("authorization token is missing")
	}

	return tokenString, nil
}
