package handlers

import (
	"backend/api/post/categoryPost"
	"backend/api/post/liquorPost"
)

// Handlers はすべてのハンドラーをまとめた構造体です。
type Handlers struct {
	LiquorHandler   *liquorPost.Handler
	CategoryHandler *categoryPost.Handler
}

// NewHandlers はHandlers構造体のコンストラクタです。
func NewHandlers(liquorHandler *liquorPost.Handler, categoryHandler *categoryPost.Handler) *Handlers {
	return &Handlers{
		LiquorHandler:   liquorHandler,
		CategoryHandler: categoryHandler,
	}
}
