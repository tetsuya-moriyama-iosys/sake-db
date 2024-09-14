package resolver

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type Resolver struct {
	CategoryRepo categoriesRepository.CategoryRepository
	LiquorRepo   liquorRepository.LiquorsRepository
	UserRepo     userRepository.UsersRepository
}

func NewResolver(categoryRepo categoriesRepository.CategoryRepository, liquorRepo liquorRepository.LiquorsRepository, userRepo userRepository.UsersRepository) *Resolver {
	return &Resolver{
		CategoryRepo: categoryRepo,
		LiquorRepo:   liquorRepo,
		UserRepo:     userRepo,
	}
}

// memo:再帰的な処理が必要であれば別途定義すること
func (r *Resolver) isFieldRequested(ctx context.Context, fieldName string) bool {
	info := graphql.GetFieldContext(ctx).Field.SelectionSet
	for _, selection := range info {
		switch field := selection.(type) {
		case *ast.Field:
			if field.Name == fieldName {
				return true
			}
		}
	}
	return false
}
