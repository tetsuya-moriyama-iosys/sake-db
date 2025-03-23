package router

import (
	"backend/di/handlers"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// ルートの設定
func graphRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	r.POST("/query", func(c *gin.Context) {
		// Ginのコンテキストからリクエストを取り出し、GraphQLの`context`にセット
		ctx := context.WithValue(c.Request.Context(), "http.Request", c.Request)
		ctx = context.WithValue(ctx, "http.ResponseWriter", c.Writer) //クッキー用
		ctx = context.WithValue(ctx, "handlers", handlers)

		// GraphQLサーバーにリクエストを渡す
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	})
	r.GET("/query", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})
}
