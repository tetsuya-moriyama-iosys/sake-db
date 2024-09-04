package categoryService

import (
	"backend/db/repository/categoriesRepository"
	"backend/graph/graphModel"
	"sort"
)

// 再帰的にカテゴリをID順にソートする関数
func sortCategories(categories []*categoriesRepository.Model) {
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].ID < categories[j].ID
	})
	for _, cat := range categories {
		if len(cat.Children) > 0 {
			sortCategories(cat.Children)
		}
	}
}

// ConvertToModelCategories 再帰的に Category を graphModel.Category に変換する
func ConvertToModelCategories(categories []*categoriesRepository.Model) []*graphModel.Category {
	var modelCategories []*graphModel.Category

	for _, category := range categories {
		// 子カテゴリを再帰的に変換
		children := ConvertToModelCategories(category.Children)

		// graphModel.Category を作成
		modelCategory := &graphModel.Category{
			ID:       category.ID,
			Name:     category.Name,
			Parent:   category.Parent, // 親カテゴリのIDをそのまま保持
			Children: children,        // 再帰的に変換された子カテゴリをセット
		}

		modelCategories = append(modelCategories, modelCategory)
	}

	return modelCategories
}

func FindCategoryByID(rootCategories []*categoriesRepository.Model, targetID int) *categoriesRepository.Model {
	for _, category := range rootCategories {
		if category.ID == targetID {
			return category
		}
		// 子カテゴリに対して再帰的に検索
		if foundCategory := FindCategoryByID(category.Children, targetID); foundCategory != nil {
			return foundCategory
		}
	}
	return nil
}
