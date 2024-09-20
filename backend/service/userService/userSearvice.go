package userService

import (
	"backend/db/repository/userRepository"
	"backend/middlewares"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUserId コンテキストからユーザーIDを取得する
func GetUserId(ctx context.Context) (*primitive.ObjectID, error) {
	userID := ctx.Value(middlewares.UserContextKey)
	if userID == nil {
		return nil, errors.New("unauthorized")
	}
	userIdObj, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		return nil, err
	}

	return &userIdObj, nil
}

func GetUserData(ctx context.Context, repo userRepository.UsersRepository) (*userRepository.Model, error) {
	userID, _ := GetUserId(ctx) //認証済みかどうかは考慮しないため空でも良いことにする
	if userID == nil {
		return nil, nil
	}
	user, err := repo.GetById(ctx, *userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 単にログイン済かどうかみたい時
func IsLogin(ctx context.Context) bool {
	userId, _ := GetUserId(ctx)
	if userId == nil {
		return false
	}
	//memo:不正なIDも一旦通す実装にする(未ログイン時とは明確に分ける必要がある)
	//if len(*userId) == 0 {
	//	return false
	//}
	return true
}
