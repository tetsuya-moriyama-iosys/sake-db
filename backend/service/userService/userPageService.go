package userService

import (
	"backend/db/repository/liquorRepository"
	"backend/graph/graphModel"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUserDetail(ctx context.Context, userId primitive.ObjectID, lRepo liquorRepository.LiquorsRepository) (*graphModel.UserDetail, error) {
	//掲示板投稿データを取得
	_, err := lRepo.BoardListByUser(ctx, userId, 10)
	if err != nil {
		return nil, err
	}

	return nil, nil

	//var comments []*graphModel.BoardPost
	//for _, comment := range boards {
	//	comments = append(comments, comment.ToGraphQL())
	//}
	//
	//return &graphModel.UserDetail{
	//	Comments:     comments,
	//	Rate5Liquors: rate5s,
	//	Rate4Liquors: rate4s,
	//	Rate3Liquors: rate3s,
	//	Rate2Liquors: rate2s,
	//	Rate1Liquors: rate1s,
	//}, nil

}
