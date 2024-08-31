package handlers

import (
	"backend/api/post/liquorPost"
)

// Handlers はすべてのハンドラーをまとめた構造体です。
type Handlers struct {
	LiquorHandler *liquorPost.Handler
}

// NewHandlers はHandlers構造体のコンストラクタです。
func NewHandlers(liquorHandler *liquorPost.Handler) *Handlers {
	return &Handlers{
		LiquorHandler: liquorHandler,
	}
}
