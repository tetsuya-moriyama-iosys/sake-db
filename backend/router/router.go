package router

import (
    "log"
    "os"
	"time"
    "github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/99designs/gqlgen/graphql/handler"

    "backend/api/auth/register"
    "backend/api/auth/login"
)


// NewRouter creates a new Gin router with GraphQL and Playground routes.
func Router(srv *handler.Server) *gin.Engine {
	// .envファイルを読み込みます
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error: loading .env file")
    }

    // 環境変数からMongoDB URIを取得します
    frontURI := os.Getenv("FRONT_URI")
    if frontURI == "" {
        log.Fatal("Error: FRONT_URI environment variable is required")
    }

    r := gin.Default()

	// CORSの設定
    config := cors.Config{
        AllowOrigins:     []string{frontURI}, // フロントエンドのURLを指定
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }

    r.Use(cors.New(config))

	// ユーザー登録とログインのエンドポイント
    r.POST("/register", func(c *gin.Context) {
        register.Register(c)
    })
    r.POST("/login", func(c *gin.Context) {
        login.Login(c)
    })

    // r.POST("/query", func(c *gin.Context) {
    //     srv.ServeHTTP(c.Writer, c.Request)
    // })
	
    // r.GET("/", func(c *gin.Context) {
    //     playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
    // })
    return r
}
