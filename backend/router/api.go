package router

import (
	"backend/di/handlers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ルートの設定
func apiRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	// 酒データの投稿
	r.POST("/post", func(c *gin.Context) {
		id, err := handlers.LiquorHandler.Post(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 正常なレスポンス
		c.JSON(http.StatusOK, gin.H{"id": *id})
	})

	// カテゴリデータの投稿
	r.POST("/category/post", func(c *gin.Context) {
		id, err := handlers.CategoryHandler.Post(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 正常なレスポンス
		c.JSON(http.StatusOK, gin.H{"id": *id})
	})
}
