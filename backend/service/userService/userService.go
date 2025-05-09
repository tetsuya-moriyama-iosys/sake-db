package userService

import (
	"backend/db/repository/userRepository"
	"backend/middlewares/auth"
	"context"
)

func GetUserData(ctx context.Context, repo userRepository.UsersRepository) (*userRepository.Model, error) {
	userID := auth.GetIdNullable(ctx) //認証済みかどうかは考慮しないため空でも良いことにする
	if userID == nil {
		return nil, nil
	}
	user, err := repo.GetById(ctx, *userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// IsLogin 単にログイン済かどうか見たい時
func IsLogin(ctx context.Context) bool {
	userId := auth.GetIdNullable(ctx)
	if userId == nil {
		return false
	}
	//memo:不正なIDも一旦通す実装にする(未ログイン時とは明確に分ける必要がある)
	//if len(*userId) == 0 {
	//	return false
	//}
	return true
}
