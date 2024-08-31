//go:build wireinject
// +build wireinject

//+ wireinject

package di

import (
	"backend/api/post/liquorPost"
	"backend/db/categoriesRepository"
	"backend/db/liquorRepository"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeHandler() (*gin.Engine, error) {
	// それぞれのnewインスタンスの生成ロジックを並べる
	wire.Build(
		BasicSet,
		// リポジトリのインスタンス生成
		categoriesRepository.NewCategoryRepository,
		liquorRepository.NewLiquorsRepository,
		//REST APIのハンドラ
		liquorPost.NewHandler,
	)
	return &gin.Engine{}, nil
}
