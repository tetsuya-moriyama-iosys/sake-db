package resolver

import (
	"backend/db/repository/bookmarkRepository"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/flavorMapRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
	DB               *mongo.Database
	CategoryRepo     categoriesRepository.CategoryRepository
	LiquorRepo       liquorRepository.LiquorsRepository
	UserRepo         userRepository.UsersRepository
	BookmarkRepo     bookmarkRepository.BookMarkRepository
	FlavorMapRepo    flavorMapRepository.FlavorMapRepository
	FlavorMapMstRepo flavorMapRepository.FlavorMapMasterRepository
	FlavorLiqRepo    flavorMapRepository.FlavorToLiquorRepository
}

func NewResolver(db *mongo.Database,
	categoryRepo categoriesRepository.CategoryRepository,
	liquorRepo liquorRepository.LiquorsRepository,
	userRepo userRepository.UsersRepository,
	bookmarkRepo bookmarkRepository.BookMarkRepository,
	flavorMapRepo flavorMapRepository.FlavorMapRepository,
	flavorMapMstRepo flavorMapRepository.FlavorMapMasterRepository,
	flavorLiqRepo flavorMapRepository.FlavorToLiquorRepository,
) *Resolver {
	return &Resolver{
		DB:               db,
		CategoryRepo:     categoryRepo,
		LiquorRepo:       liquorRepo,
		UserRepo:         userRepo,
		BookmarkRepo:     bookmarkRepo,
		FlavorMapRepo:    flavorMapRepo,
		FlavorMapMstRepo: flavorMapMstRepo,
		FlavorLiqRepo:    flavorLiqRepo,
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
