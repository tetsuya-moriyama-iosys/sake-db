package router

import (
	"backend/api/sso/x"
	"backend/di/handlers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ルートの設定
func oauthRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	// Xログイン用のエンドポイント
	r.GET("/x/login", func(c *gin.Context) {
		url, err := x.GenerateAuthURL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
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
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
}
