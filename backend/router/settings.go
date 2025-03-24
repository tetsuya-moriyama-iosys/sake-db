package router

import (
	"backend/di/handlers"
	"backend/middlewares"
	"backend/middlewares/customError/logger"
	"backend/util/helper"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func Router(srv *handler.Server, handlers *handlers.Handlers) *gin.Engine {
	// .envファイルを読み込みます
	helper.LoadEnv()

	r := gin.Default()

	// CORS設定
	//r.Use(corsMiddleware(getFrontURI()))

	// エラーハンドリング
	logger.Init(*handlers.ErrorHandler)
	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.GinCustomRecovery())

	// GraphQL のエラーハンドリング設定
	srv.SetErrorPresenter(middlewares.GraphQLErrorPresenter) // GraphQL のエラーを適切にログ出力
	srv.SetRecoverFunc(middlewares.GraphQLRecover)           // panic からの復旧

	// ルート設定
	configureRoutes(r, srv, handlers)

	return r
}

// FRONT_URI環境変数の取得
//func getFrontURI() string {
//	frontURI := os.Getenv("FRONT_URI")
//	if frontURI == "" {
//		log.Fatal("Error: FRONT_URI environment variable is required")
//	}
//	return frontURI
//}

// CORSミドルウェアの設定
//func corsMiddleware(frontURI string) gin.HandlerFunc {
//	config := cors.Config{
//		AllowOrigins:     []string{frontURI},
//		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
//		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
//		ExposeHeaders:    []string{"Content-Length"},
//		AllowCredentials: true, // クッキーを有効にする
//		MaxAge:           12 * time.Hour,
//	}
//	return cors.New(config)
//}
