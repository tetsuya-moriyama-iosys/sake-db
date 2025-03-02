package router

import (
	"backend/di/handlers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

// ルートの設定
func configureRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	apiRoutes(r, srv, handlers)
	oauthRoutes(r, srv, handlers)
	graphRoutes(r, srv, handlers)
}
