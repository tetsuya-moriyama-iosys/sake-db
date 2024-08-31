package router

import (
	"backend/di/handlers"
	"backend/util/helper"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func Router(srv *handler.Server, handlers *handlers.Handlers) *gin.Engine {
	// .envファイルを読み込みます
	helper.LoadEnv()

	// フロントエンドのURIを取得
	frontURI := getFrontURI()

	r := gin.Default()

	// CORS設定
	r.Use(corsMiddleware(frontURI))

	/// ルート設定
	configureRoutes(r, srv, handlers)

	return r
}

// FRONT_URI環境変数の取得
func getFrontURI() string {
	frontURI := os.Getenv("FRONT_URI")
	if frontURI == "" {
		log.Fatal("Error: FRONT_URI environment variable is required")
	}
	return frontURI
}

// CORSミドルウェアの設定
func corsMiddleware(frontURI string) gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     []string{frontURI},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(config)
}
