package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strings"
)

// Claims represents the JWT claims
type Claims struct {
	Id primitive.ObjectID `json:"id"` //これがJWTトークンに含まれる
	jwt.RegisteredClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
var (
	ErrInvalidToken = errors.New("unauthorized")
	ErrTokenExpired = errors.New("token expired")
)

// ContextKey ユーザーIDを保存するためのコンテキストキー
type ContextKey string

const UserContextKey ContextKey = "user"

// AuthenticateJWT JWTの認証ミドルウェア TODO: REST用だが現状使われていなさそう。おそらくログ残す際に必要。
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// "Bearer "プレフィックスを除去してトークンを取得
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}
		// トークンの解析
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 署名方法が正しいかチェック
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JwtKey, nil
		})

		// トークンが無効または解析に失敗した場合
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// ユーザーIDをコンテキストに保存
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := claims["user_id"].(string)
			ctx := context.WithValue(c.Request.Context(), UserContextKey, userID)
			c.Request = c.Request.WithContext(ctx)
		}

		c.Set("userId", claims.Id)
		c.Next()
	}
}

// ExtractTokenFromHeader JWTトークンを読み込むための関数
func ExtractTokenFromHeader(ctx context.Context) (string, error) {
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

func AuthenticateToken(ctx context.Context, tokenString string) (context.Context, error) {
	// トークンのパースと検証
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtKey, nil
	})

	if err != nil {
		// トークンのエラーを分類
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			// トークンが期限切れの場合
			if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
				return ctx, ErrTokenExpired
			}
		}
		return ctx, ErrInvalidToken
	}

	if !token.Valid {
		return ctx, ErrInvalidToken
	}

	// 認証に成功した場合、ユーザーIDをcontextに保存
	ctx = context.WithValue(ctx, UserContextKey, claims.Id)

	return ctx, nil
}
