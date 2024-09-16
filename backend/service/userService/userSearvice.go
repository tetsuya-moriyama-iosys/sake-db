package userService

import (
	"backend/db/repository/userRepository"
	"backend/middlewares"
	"context"
	"errors"
)

// GetUserId コンテキストからユーザーIDを取得する
func GetUserId(ctx context.Context) (*string, error) {
	userID := ctx.Value(middlewares.UserContextKey)
	if userID == nil {
		return nil, errors.New("unauthorized")
	}
	userIdStr := userID.(string)

	return &userIdStr, nil
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
