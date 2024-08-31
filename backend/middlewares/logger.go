package middlewares

// Logger はリクエストとコンテキストのログを出力するミドルウェアです。
//func Logger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		log.Println("Middleware called")
//
//		if c == nil {
//			log.Println("Error: *gin.Context is nil in middleware")
//		} else {
//			log.Println("Middleware context is not nil")
//		}
//
//		// 次のハンドラーに処理を渡す
//		c.Next()
//	}
//}
