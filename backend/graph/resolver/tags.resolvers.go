package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"backend/db/repository/liquorRepository"
	"backend/graph/graphModel"
	"backend/service/userService"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostTag is the resolver for the postTag field.
func (r *mutationResolver) PostTag(ctx context.Context, input graphModel.TagInput) (*graphModel.Tag, error) {
	uId, err := userService.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	lId, err := primitive.ObjectIDFromHex(input.LiquorID)
	tag, err := r.LiquorRepo.PostTag(ctx, lId, uId, input.Text)
	if err != nil {
		return nil, err
	}
	return tag.ToGraphQL(), nil
}

// DeleteTag is the resolver for the deleteTag field.
func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	tId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	err = r.LiquorRepo.DeleteTag(ctx, tId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetTags is the resolver for the getTags field.
func (r *queryResolver) GetTags(ctx context.Context, liquorID string) ([]*graphModel.Tag, error) {
	tId, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, err
	}
	tags, err := r.LiquorRepo.GetTags(ctx, tId)
	if err != nil {
		return nil, err
	}
	return liquorRepository.TagsToGraphQL(tags), nil
}
