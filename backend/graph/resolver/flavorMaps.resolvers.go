package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/graph/graphModel"
	"backend/middlewares/auth"
	"backend/service/flavorMapService"
	"backend/util/utilType"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostFlavor is the resolver for the postFlavor field.
func (r *mutationResolver) PostFlavor(ctx context.Context, input graphModel.PostFlavorMap) (bool, error) {
	lId, err := primitive.ObjectIDFromHex(input.LiquorID)
	if err != nil {
		return false, err
	}

	//マスタが存在するのを確認したので、フレーバーマップを更新する
	err = flavorMapService.PostFlavorMap(ctx, &r.FlavorMapMstRepo, &r.FlavorLiqRepo, &r.FlavorMapRepo, &r.CategoryRepo, &r.LiquorRepo, lId, utilType.Coordinates{
		X: input.X,
		Y: input.Y,
	})
	return err == nil, err
}

// GetFlavorMap is the resolver for the getFlavorMap field.
func (r *queryResolver) GetFlavorMap(ctx context.Context, liquorID string) (*graphModel.FlavorMapData, error) {
	lId, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, err
	}
	result, err := flavorMapService.GetFlavorMap(ctx, &r.FlavorMapMstRepo, &r.FlavorLiqRepo, &r.LiquorRepo, &r.CategoryRepo, lId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		//フレーバーマップがない場合
		return nil, nil
	}

	return result.ToGraphQL(), nil
}

// GetVoted is the resolver for the getVoted field.
func (r *queryResolver) GetVoted(ctx context.Context, liquorID string) (*graphModel.VotedData, error) {
	uId, err := auth.GetId(ctx)
	if err != nil {
		return nil, err
	}
	lId, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, err
	}
	mst, err := flavorMapService.GetFlavorMasterData(ctx, &r.FlavorMapMstRepo, &r.LiquorRepo, &r.CategoryRepo, lId)
	if err != nil {
		return nil, err
	}
	result, err := r.FlavorMapRepo.GetVotedDataByLiquor(ctx, uId, lId, mst.CategoryID)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result.ToGraphQL(), nil
}
