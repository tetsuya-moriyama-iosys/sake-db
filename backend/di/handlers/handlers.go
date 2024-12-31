package handlers

import (
	"backend/api"
	"backend/api/post/categoryPost"
	"backend/api/post/liquorPost"
)

// Handlers はすべてのハンドラーをまとめた構造体です。
type Handlers struct {
	LiquorHandler   *liquorPost.Handler
	CategoryHandler *categoryPost.Handler
	UserHandler     *api.UserHandler
}

// NewHandlers はHandlers構造体のコンストラクタです。
func NewHandlers(liquorHandler *liquorPost.Handler, categoryHandler *categoryPost.Handler, userHandler *api.UserHandler) *Handlers {
	return &Handlers{
		LiquorHandler:   liquorHandler,
		CategoryHandler: categoryHandler,
		UserHandler:     userHandler,
	}
}
