package router

import (
	"backend/di/handlers"
	"backend/middlewares/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ルートの設定
func apiRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	// 任意認証が必要
	// 酒データの投稿
	r.POST("/post", auth.RESTOptionalAuthenticate(handlers.TokenConfig), func(c *gin.Context) {
		id, err := handlers.LiquorHandler.Post(c, &handlers.UserHandler.UserRepo)
		if err != nil {
			_ = c.Error(err)
			return
		}
		// 正常なレスポンス
		c.JSON(http.StatusOK, gin.H{"id": *id})
	})

	// カテゴリデータの投稿
	r.POST("/category/post", auth.RESTOptionalAuthenticate(handlers.TokenConfig), func(c *gin.Context) {
		id, err := handlers.CategoryHandler.Post(c, &handlers.UserHandler.UserRepo)
		if err != nil {
			_ = c.Error(err)
			return
		}
		// 正常なレスポンス
		c.JSON(http.StatusOK, gin.H{"id": *id})
	})
}
