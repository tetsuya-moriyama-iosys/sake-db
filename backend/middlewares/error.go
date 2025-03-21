package middlewares

import (
	"backend/customError"
	"backend/customError/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler はエラーをキャッチし、適切なレスポンスを返すミドルウェア
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Ginのエラーリストからカスタムエラーを取得
			for _, e := range c.Errors {
				var customErr *customError.Error
				if errors.As(e.Err, &customErr) {
					// エラーログを記録
					logger.LogError(customErr)

					// クライアントにはユーザーフレンドリーなエラーメッセージを返す
					c.JSON(customErr.StatusCode, gin.H{
						"error_id": customErr.ID,
						"message":  customErr.UserMessage,
					})
					return
				}
			}

			// 予期しないエラーの場合
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_id": "unknown",
				"message":  "Internal Server Error",
			})
		}
	}
}
