package resolver

import (
	"backend/di/handlers"
	"backend/service/authService/tokenConfig"
	"context"
	"net/http"
)

func getHttpRequest(c context.Context) *http.Request {
	// context.Contextからhttp.Requestを取得
	req, ok := c.Value("http.Request").(*http.Request)
	if !ok {
		panic("failed to retrieve *http.Request from context")
	}
	return req
}

func getResponseWriter(c context.Context) http.ResponseWriter {
	// context.Contextからhttp.ResponseWriterを取得
	writer, ok := c.Value("http.ResponseWriter").(http.ResponseWriter)
	if !ok {
		panic("http.ResponseWriterがcontextに存在しません")
	}
	return writer
}

func getHandler(c context.Context) *handlers.Handlers {
	h, ok := c.Value("handlers").(*handlers.Handlers)
	if !ok {
		panic("handlerが存在しません")
	}
	return h
}

// GetTokenConfig directiveから呼び出す必要があったため公開した
func GetTokenConfig(c context.Context) *tokenConfig.TokenConfig {
	h := getHandler(c)
	return h.TokenConfig
}
