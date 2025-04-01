package middlewares

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
	"time"
)

// ErrorHandler はエラーをキャッチし、適切なレスポンスを返すミドルウェア
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		var customErr *customError.Error
		if errors.As(err.Err, &customErr) {
			// エラーログを記録
			logger.LogError(c, customErr)

			// クライアントにはユーザーフレンドリーなエラーメッセージを返す
			c.JSON(customErr.StatusCode, gin.H{
				"error_id": customErr.ID,
				"message":  customErr.UserMessage,
			})
			return
		}
		// カスタムエラーを定義していないエラー
		// エラーログを記録
		pc, file, line, ok := runtime.Caller(2) //一旦2で類推している
		if !ok {
			return
		}
		fn := runtime.FuncForPC(pc)
		customErr = &customError.Error{
			ID:          uuid.New().String(),
			StatusCode:  http.StatusInternalServerError,
			ErrorCode:   "unknown",
			UserMessage: "エラーが発生しました。",
			Level:       logrus.ErrorLevel, // 未定義エラーなので深刻なエラーとして扱う
			RawErr:      err.Err,
			Location:    fmt.Sprintf("%s:%d (%s)", file, line, fn.Name()),
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		logger.LogError(c, customErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Err.Error(),
		})
	}
}

// GinCustomRecovery は CustomRecoveryWithWriter を返すラッパー
func GinCustomRecovery() gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, recovered interface{}) {
		err := logger.LogPanic(c.Request.Context(), recovered, "Gin")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.UserMessage,
		})
		c.Abort()
	})
}
