package categoryService

import (
	"backend/db/categoriesRepository"
	"backend/graph/model"
	"context"
)

// LeveledCategoriesGet 階層分けされたカテゴリを取得する
func LeveledCategoriesGet(ctx context.Context, r *categoriesRepository.CategoryRepository) ([]*model.Category, error) {
	//DBからデータを取得
	categories, err := r.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	//ID順にソートする(TODO:DBにやらせた方がいい気がするが、後で調べて実装する)
	sortCategories(categories)

	modelCatList := make([]*model.Category, len(categories))
	for i, cat := range categories {
		modelCatList[i] = ConvertToModelCategory(cat)
	}
	return modelCatList, nil
}
