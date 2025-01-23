package router

import (
	"backend/di/handlers"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ルートの設定
func graphRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
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
		// Ginのコンテキストからリクエストを取り出し、GraphQLの`context`にセット
		ctx := context.WithValue(c.Request.Context(), "http.Request", c.Request)

		// GraphQLサーバーにリクエストを渡す
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	})
	r.GET("/query", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})
}
