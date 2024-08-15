package router

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"backend/api/auth/login"
	"backend/api/auth/register"
	"github.com/99designs/gqlgen/graphql/playground"
)

//// カスタムエラープレゼンター
//func customErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
//	// ログ出力
//	log.Printf("GraphQL Error: %v", err)
//
//	// もともとのエラーをそのまま返す
//	return graphql.DefaultErrorPresenter(ctx, err)
//}

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

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Printf("GraphQL Error: %v", err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})

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

	// GraphQLインターフェース
	r.POST("/query", func(c *gin.Context) {

		// リクエストボディを読み取り、ログに出力
		//_, err := io.ReadAll(c.Request.Body)
		//if err != nil {
		//	log.Printf("Error reading request body: %v", err)
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot read request body"})
		//	return
		//}
		//log.Printf("c.Writer: %s", c.Writer)
		//log.Println("c.Request:", c.Request)

		// 読み取ったボディを元に戻す
		//c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/query", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	return r
}
