package auth

import (
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

// RESTAuthenticate REST用のもの(Optionalしか使わないかも)
//func RESTAuthenticate(tokenConfig *tokenConfig.TokenConfig) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString, err := ExtractTokenFromHeader(c.Request)
//		if err != nil {
//			_ = c.Error(err)
//			c.Abort()
//			return
//		}
//		_, err = AuthenticateToken(c, tokenString, *tokenConfig)
//		if err != nil {
//			_ = c.Error(err)
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}

func RESTOptionalAuthenticate(tokenConfig *tokenConfig.TokenConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := ExtractTokenFromHeader(c.Request)
		if err != nil {
			// トークンが存在しない場合はそのまま処理を続行
			c.Next()
			return
		}
		_, err = AuthenticateToken(c, tokenString, *tokenConfig)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

// ExtractTokenFromHeader JWTトークンを読み込むための関数
func ExtractTokenFromHeader(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", errMissHeader()
	}

	// "Bearer "を除去してトークンを取得
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return "", errMissBearer()
	}

	return tokenString, nil
}

func AuthenticateToken(ctx context.Context, tokenString string, tokenConfig tokenConfig.TokenConfig) (context.Context, error) {
	// トークンのパースと検証
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenConfig.AccessSecretKey, nil
	})

	if err == nil {
		if !token.Valid {
			return ctx, errTokenInvalid(err)
		}
		// 認証に成功した場合、ユーザーIDをcontextに保存
		ctx = setId(ctx, claims.Id)
		return ctx, nil
	}

	// 認証に失敗
	var validationErr *jwt.ValidationError
	if errors.As(err, &validationErr) {
		// トークンが期限切れの場合
		if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
			return ctx, errTokenExpired(err)
		}
	}
	return ctx, errTokenSomething(err)
}
