package middlewares

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

// Claims represents the JWT claims
type Claims struct {
	Id string `json:"id"` //これがJWTトークンに含まれる
	jwt.RegisteredClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// ContextKey ユーザーIDを保存するためのコンテキストキー
type ContextKey string

const UserContextKey ContextKey = "user" //TODO:キーの内容が合ってるかどうかは後で確認すること

// AuthenticateJWT JWTの認証ミドルウェア
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
