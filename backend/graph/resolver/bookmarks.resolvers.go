package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"backend/db/repository/bookmarkRepository"
	"backend/graph/graphModel"
	"backend/middlewares/auth"
	"backend/service/bookmarkService"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddBookMark is the resolver for the addBookMark field.
func (r *mutationResolver) AddBookMark(ctx context.Context, id string) (bool, error) {
	uId, targetId, err := bookmarkService.GetPrimitiveIds(ctx, id)
	if err != nil {
		return false, err
	}
	err = r.BookmarkRepo.Add(ctx, uId, targetId)
	if err != nil {
		return false, err
	}
	return true, err
}

// RemoveBookMark is the resolver for the removeBookMark field.
func (r *mutationResolver) RemoveBookMark(ctx context.Context, id string) (bool, error) {
	uId, targetId, err := bookmarkService.GetPrimitiveIds(ctx, id)
	if err != nil {
		return false, err
	}
	err = r.BookmarkRepo.Remove(ctx, uId, targetId)
	if err != nil {
		return false, err
	}
	return true, err
}

// GetIsBookMarked is the resolver for the getIsBookMarked field.
func (r *queryResolver) GetIsBookMarked(ctx context.Context, id string) (bool, error) {
	uId, targetId, err := bookmarkService.GetPrimitiveIds(ctx, id)
	if err != nil {
		return false, err
	}
	_, err = r.BookmarkRepo.Find(ctx, uId, targetId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// ドキュメントが存在しないエラーは、単にfalseを返せばいい
			return false, nil
		}
		//それ以外のエラーは普通にエラー
		return false, err
	}
	return true, nil
}

// GetRecommendLiquorList is the resolver for the getRecommendLiquorList field.
func (r *queryResolver) GetRecommendLiquorList(ctx context.Context) ([]*graphModel.Recommend, error) {
	uId, err := auth.GetId(ctx)
	if err != nil {
		return nil, err
	}
	list, err := r.BookmarkRepo.GetRecommendLiquors(ctx, uId, nil)
	if err != nil {
		return nil, err
	}

	return list.ToGraphQL(), nil
}

// GetBookMarkList is the resolver for the getBookMarkList field.
func (r *queryResolver) GetBookMarkList(ctx context.Context) ([]*graphModel.BookMarkListUser, error) {
	uId, err := auth.GetId(ctx)
	if err != nil {
		return nil, err
	}
	bList, err := r.BookmarkRepo.List(ctx, uId)
	if err != nil {
		return nil, err
	}
	return bookmarkRepository.BookMarkList(bList).ToGraphQL(), nil
}

// GetBookMarkedList is the resolver for the getBookMarkedList field.
func (r *queryResolver) GetBookMarkedList(ctx context.Context, id string) ([]*graphModel.BookMarkListUser, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	bList, err := r.BookmarkRepo.BookmarkedList(ctx, idObj)
	if err != nil {
		return nil, err
	}
	return bookmarkRepository.BookMarkList(bList).ToGraphQL(), nil
}
