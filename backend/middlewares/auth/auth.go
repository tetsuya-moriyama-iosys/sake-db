package auth

import (
	"backend/di/handlers"
	"backend/service/authService/tokenConfig"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

// Claims represents the JWT claims
type Claims struct {
	Id primitive.ObjectID `json:"id"` //これがJWTトークンに含まれる
	jwt.RegisteredClaims
}

var (
	ErrInvalidToken = errors.New("unauthorized")
	ErrTokenExpired = errors.New("token expired")
)

// RESTAuthenticate TODO: REST用
func RESTAuthenticate(handlers *handlers.Handlers) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := ExtractTokenFromHeader(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		_, err = AuthenticateToken(c, tokenString, *handlers.TokenConfig)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

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

func AuthenticateToken(ctx context.Context, tokenString string, tokenConfig tokenConfig.TokenConfig) (context.Context, error) {
	// トークンのパースと検証
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenConfig.AccessSecretKey, nil
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
	ctx = setId(ctx, claims.Id)

	return ctx, nil
}
