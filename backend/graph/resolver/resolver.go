package resolver

import (
	"backend/db/categoriesRepository"
	"backend/db/liquorRepository"
)

type Resolver struct {
	CategoryRepo categoriesRepository.CategoryRepository
	LiquorRepo   liquorRepository.LiquorsRepository
}

func NewResolver(categoryRepo categoriesRepository.CategoryRepository, liquorRepo liquorRepository.LiquorsRepository) *Resolver {
	return &Resolver{
		CategoryRepo: categoryRepo,
		LiquorRepo:   liquorRepo,
	}
}
