package router

import (
	"backend/api/post/liquor"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"backend/api/auth/login"
	"backend/api/auth/register"
	"github.com/99designs/gqlgen/graphql/playground"
)

//func ErrorHandlingMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Next() // リクエストを次のハンドラーに渡す// ミドルウェア後にエラーがあるかどうかを確認iflen(c.Errors) > 0 {
//		for _, e := range c.Errors {
//			log.Printf("Error occurred: %v", e.Err)
//		}
//		// 必要に応じて適切なレスポンスを返す
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
//	}
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

	//srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
	//	log.Printf("GraphQL Error: %v", err)
	//	return graphql.DefaultErrorPresenter(ctx, err)
	//})

	r := gin.Default()

	// エラーハンドリングミドルウェアの追加
	//r.Use(ErrorHandlingMiddleware())

	// CORSの設定
	config := cors.Config{
		AllowOrigins:     []string{frontURI}, // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
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
	//投稿のエンドポイント
	r.POST("/post", func(c *gin.Context) {
		liquor.Post(c)
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
		// リクエストボディをログに出力
		//bodyBytes, err := io.ReadAll(c.Request.Body)
		//if err != nil {
		//	log.Printf("Error reading request body: %v", err)
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		//	return
		//}
		//log.Printf("Request Body: %s", string(bodyBytes))

		// 読み取ったボディを元に戻す
		//c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// GraphQLサーバーにリクエストを渡す
		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/query", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	return r
}
