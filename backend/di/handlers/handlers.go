package handlers

import (
	"backend/api/post/categoryPost"
	"backend/api/post/liquorPost"
	"backend/service/authService/tokenConfig"
)

// Handlers はすべてのハンドラーをまとめた構造体です。
type Handlers struct {
	LiquorHandler   *liquorPost.Handler
	CategoryHandler *categoryPost.Handler
	TokenConfig     *tokenConfig.TokenConfig
}

// NewHandlers はHandlers構造体のコンストラクタです。
func NewHandlers(liquorHandler *liquorPost.Handler, categoryHandler *categoryPost.Handler, tokenConfig *tokenConfig.TokenConfig) *Handlers {
	return &Handlers{
		LiquorHandler:   liquorHandler,
		CategoryHandler: categoryHandler,
		TokenConfig:     tokenConfig,
	}
}
