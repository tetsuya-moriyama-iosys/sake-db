package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/graph/graphModel"
	"backend/service/categoryService"
	"context"
)

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*graphModel.Category, error) {
	// 構造化されたカテゴリ一覧を返す
	categories, err := categoryService.LeveledCategoriesGet(ctx, &r.CategoryRepo)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
