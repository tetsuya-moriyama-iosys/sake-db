package router

import (
	"backend/di/handlers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ルートの設定
func configureRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	log.Println("in router")

	//r.POST("/register", func(c *gin.Context) {
	//	register.Register(c)
	//})
	//r.POST("/login", func(c *gin.Context) {
	//	login.Login(c)
	//})

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

	// GraphQLインターフェース
	r.POST("/query", func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic occurred: %v", r)
				if !c.Writer.Written() { // レスポンスがまだ書き込まれていない場合
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					log.Println("Response already written, skipping error response")
				}
			}
		}()
		// GraphQLサーバーにリクエストを渡す
		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/query", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})
}
