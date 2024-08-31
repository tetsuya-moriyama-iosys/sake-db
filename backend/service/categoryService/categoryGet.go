package categoryService

import (
	"backend/db/categoriesRepository"
	"context"
)

// LeveledCategoriesGet 階層分けされたカテゴリを取得する
func LeveledCategoriesGet(ctx context.Context, r *categoriesRepository.CategoryRepository) ([]*categoriesRepository.Category, error) {
	//DBからデータを取得
	categories, err := r.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	//ID順にソートする(TODO:DBにやらせた方がいい気がするが、後で調べて実装する)
	sortCategories(categories)

	// カテゴリをIDをキーとするマップに格納
	categoryMap := make(map[int]*categoriesRepository.Category)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}
	// 親子関係を構築
	var rootCategories []*categoriesRepository.Category
	for _, cat := range categories {
		if cat.Parent != nil {
			// Parent が存在する場合、親カテゴリの Children に追加
			parentCategory, exists := categoryMap[*cat.Parent]
			if exists {
				parentCategory.Children = append(parentCategory.Children, cat)
			}
		} else {
			// Parent が存在しない場合、ルートカテゴリとして扱う
			rootCategories = append(rootCategories, cat)
		}
	}

	return rootCategories, nil
}
