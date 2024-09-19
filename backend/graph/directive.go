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

// トークンを検証し、ユーザーIDをcontextに保存する共通関数
func authenticateToken(ctx context.Context, tokenString string) (context.Context, error) {
	// トークンのパースと検証
	claims := &middlewares.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return middlewares.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return ctx, errors.New("invalid token")
	}

	// 認証に成功した場合、ユーザーIDをcontextに保存
	ctx = context.WithValue(ctx, middlewares.UserContextKey, claims.Id)

	return ctx, nil
}

// 必須認証のディレクティブ
func authDirective(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	// ヘッダーからトークンを取得
	tokenString, err := extractTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// トークンを検証し、ユーザーIDをcontextに保存
	ctx, err = authenticateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	// 認証に成功した場合、次のリゾルバを実行
	return next(ctx)
}

// 任意認証のディレクティブ
func optionalAuthDirective(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	// ヘッダーからトークンを取得
	tokenString, err := extractTokenFromHeader(ctx)

	// トークンが存在しない場合は、認証なしで処理を続行
	if err != nil {
		return next(ctx)
	}

	// トークンを検証し、認証が成功すればユーザーIDをcontextに保存
	ctx, err = authenticateToken(ctx, tokenString)
	if err != nil {
		// トークンが無効ならエラーを返す
		return nil, err
	}

	// 認証に成功した場合、次のリゾルバを実行
	return next(ctx)
}
