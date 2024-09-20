package userService

import (
	"backend/db/repository/liquorRepository"
	"backend/graph/graphModel"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUserDetail(ctx context.Context, userId primitive.ObjectID, lRepo liquorRepository.LiquorsRepository) (*graphModel.UserDetail, error) {
	//掲示板投稿データを取得
	boards, err := lRepo.BoardListByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	var rate5Ids []primitive.ObjectID
	var rate4Ids []primitive.ObjectID
	var rate3Ids []primitive.ObjectID
	var rate2Ids []primitive.ObjectID
	var rate1Ids []primitive.ObjectID

	// 各ボードをループして、rateフィールドを確認
	for _, board := range boards {
		if board.Rate == nil {
			// rateがnilの場合はスキップ
			continue
		}

		// rateの値に応じて対応する配列にLiquorIDを追加
		switch *board.Rate {
		case 5:
			rate5Ids = append(rate5Ids, board.LiquorID)
		case 4:
			rate4Ids = append(rate4Ids, board.LiquorID)
		case 3:
			rate3Ids = append(rate3Ids, board.LiquorID)
		case 2:
			rate2Ids = append(rate2Ids, board.LiquorID)
		case 1:
			rate1Ids = append(rate1Ids, board.LiquorID)
		}
	}

	rate5Liquors, err := lRepo.GetLiquorsByIds(ctx, rate5Ids)
	if err != nil {
		return nil, err
	}
	rate4Liquors, err := lRepo.GetLiquorsByIds(ctx, rate4Ids)
	if err != nil {
		return nil, err
	}
	rate3Liquors, err := lRepo.GetLiquorsByIds(ctx, rate3Ids)
	if err != nil {
		return nil, err
	}
	rate2Liquors, err := lRepo.GetLiquorsByIds(ctx, rate2Ids)
	if err != nil {
		return nil, err
	}
	rate1Liquors, err := lRepo.GetLiquorsByIds(ctx, rate1Ids)
	if err != nil {
		return nil, err
	}
	var rate5s []*graphModel.LiquorSimple
	var rate4s []*graphModel.LiquorSimple
	var rate3s []*graphModel.LiquorSimple
	var rate2s []*graphModel.LiquorSimple
	var rate1s []*graphModel.LiquorSimple

	for _, liquor := range rate5Liquors {
		rate5s = append(rate5s, liquor.ToGraphQLSimple())
	}
	for _, liquor := range rate4Liquors {
		rate4s = append(rate4s, liquor.ToGraphQLSimple())
	}
	for _, liquor := range rate3Liquors {
		rate3s = append(rate3s, liquor.ToGraphQLSimple())
	}
	for _, liquor := range rate2Liquors {
		rate2s = append(rate2s, liquor.ToGraphQLSimple())
	}
	for _, liquor := range rate1Liquors {
		rate1s = append(rate1s, liquor.ToGraphQLSimple())
	}

	var comments []*graphModel.BoardPost
	for _, comment := range boards {
		comments = append(comments, comment.ToGraphQL())
	}

	return &graphModel.UserDetail{
		Comments:     comments,
		Rate5Liquors: rate5s,
		Rate4Liquors: rate4s,
		Rate3Liquors: rate3s,
		Rate2Liquors: rate2s,
		Rate1Liquors: rate1s,
	}, nil

}
