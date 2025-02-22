package router

import (
	"backend/api/sso/x"
	"backend/di/handlers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// ルートの設定
func oauthRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	// Xログイン用のエンドポイント
	r.GET("/x/login", func(c *gin.Context) {
		url, err := x.GenerateAuthURL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return //TODO:エラーページに飛ばす
		}
		// JSONレスポンスでURLを返す
		c.JSON(http.StatusOK, gin.H{
			"redirectUrl": url,
		})
	})

	r.GET("/x/callback", func(c *gin.Context) {
		user, err := x.Login(handlers, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return //TODO:エラーページに飛ばす
		}

		//一時的なクッキーを作る
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     os.Getenv("SSO_AUTH_TOKEN_NAME"),
			Value:    user.Token,
			Path:     "/",
			Domain:   "",
			MaxAge:   60,
			Secure:   false, //TODO:HTTPS対応が済んだらtrueにする
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		// フロントエンドにリダイレクト
		c.Redirect(http.StatusFound, "http://frontend-app-url/")
		c.JSON(http.StatusOK, gin.H{"user": user})

		//TODO:JWTトークンのみをクッキーに保存し、認証はストアからJWTトークンを用いた認証を改めて行う(userデータをクッキーに含めたくないため)
	})
}
