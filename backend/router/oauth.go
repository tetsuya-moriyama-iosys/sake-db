package router

import (
	"backend/api/sso/x"
	"backend/di/handlers"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// ルートの設定
func oauthRoutes(r *gin.Engine, srv *handler.Server, handlers *handlers.Handlers) {
	// Xログイン用のエンドポイント
	r.GET("/x/login", func(c *gin.Context) {
		url, err := x.GenerateAuthURL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return //TODO:エラーページに飛ばす
		}
		// JSONレスポンスでURLを返す
		c.JSON(http.StatusOK, gin.H{
			"redirectUrl": url,
		})
	})

	r.GET("/x/callback", func(c *gin.Context) {
		// Ginのコンテキストからリクエストを取り出し、GraphQLの`context`にセット(GraphQL側の処理と共通化してるため、合わせる)
		ctx := context.WithValue(c.Request.Context(), "http.Request", c.Request)
		ctx = context.WithValue(ctx, "http.ResponseWriter", c.Writer) //クッキー用

		user, err := x.Login(c, handlers, c.Writer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return //TODO:エラーページに飛ばす
		}

		// フロントエンドにリダイレクト
		frontURI := os.Getenv("FRONT_URI")
		c.Redirect(http.StatusFound, frontURI)
		c.JSON(http.StatusFound, gin.H{"user": user})
	})
}
