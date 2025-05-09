//go:build wireinject
// +build wireinject

//+ wireinject

package di

import (
	"backend/api"
	"backend/api/post/categoryPost"
	"backend/api/post/liquorPost"
	"backend/db/repository/bookmarkRepository"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/errorRepository"
	"backend/db/repository/flavorMapRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"backend/service/authService/tokenConfig"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeHandler() (*gin.Engine, error) {
	// それぞれのnewインスタンスの生成ロジックを並べる
	wire.Build(
		tokenConfig.NewTokenConfig,
		BasicSet,
		//REST APIのハンドラ
		liquorPost.NewHandler,
		categoryPost.NewHandler,
		api.NewUserHandler,
		// リポジトリのインスタンス生成
		categoriesRepository.NewCategoryRepository,
		liquorRepository.NewLiquorsRepository,
		userRepository.NewUsersRepository,
		bookmarkRepository.NewBookMarkRepository,
		flavorMapRepository.NewFlavorMapMasterRepository,
		flavorMapRepository.NewFlavorMapRepository,
		flavorMapRepository.NewFlavorToLiquorRepository,
		errorRepository.New,
	)
	return &gin.Engine{}, nil
}
