package userService

import (
	"backend/db/repository/liquorRepository"
	"backend/middlewares/customError"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUserDetail(ctx context.Context, userId primitive.ObjectID, lRepo liquorRepository.LiquorsRepository) (*liquorRepository.BoardListResponse, *customError.Error) {
	//掲示板投稿データを取得
	boards, err := lRepo.BoardListByUser(ctx, userId, 10)
	if err != nil {
		return nil, err
	}
	return boards, nil
}
