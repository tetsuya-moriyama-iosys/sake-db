package categoryService

import (
	"backend/db/repository/categoriesRepository"
	"backend/middlewares/customError"
	"context"
)

// LeveledCategoriesGet 階層分けされたカテゴリを取得する
func LeveledCategoriesGet(ctx context.Context, r *categoriesRepository.CategoryRepository) ([]*categoriesRepository.Model, *customError.Error) {
	//DBからデータを取得
	categories, err := r.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	//ID順にソートする(TODO:DBにやらせた方がいい気がするが、後で調べて実装する)
	sortCategories(categories)

	// カテゴリをIDをキーとするマップに格納
	categoryMap := make(map[int]*categoriesRepository.Model)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}
	// 親子関係を構築
	var rootCategories []*categoriesRepository.Model
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

func PartialLeveledCategoriesGet(ctx context.Context, targetId int, r *categoriesRepository.CategoryRepository) (*categoriesRepository.Model, *customError.Error) {
	//DBからデータを全件取得
	categories, err := LeveledCategoriesGet(ctx, r)
	if err != nil {
		return nil, err
	}
	for _, category := range categories {
		if category.ID == targetId {
			return category, nil
		}
		// 子カテゴリに対して再帰的に検索
		if foundCategory := FindCategoryByID(category.Children, targetId); foundCategory != nil {
			return foundCategory, nil
		}
	}
	return nil, nil
}

func GetBelongCategoryIdList(ctx context.Context, targetId int, r *categoriesRepository.CategoryRepository) ([]int, *customError.Error) {
	var result []int
	categoryList, err := PartialLeveledCategoriesGet(ctx, targetId, r)
	if err != nil {
		return result, err
	}

	// IDリストを取得するためのヘルパー関数
	var collectCategoryIDs func(category *categoriesRepository.Model)
	collectCategoryIDs = func(category *categoriesRepository.Model) {
		// まず、自身のIDを追加
		result = append(result, category.ID)
		for _, child := range category.Children {
			// 子カテゴリがある場合は再帰的にIDを収集
			collectCategoryIDs(child)
		}
	}

	// categoryListのIDを収集
	collectCategoryIDs(categoryList)

	return result, nil
}

// GetCategoryTrail 指定されたカテゴリIDのパンくずリストを配列として作成する
func GetCategoryTrail(ctx context.Context, targetId int, r *categoriesRepository.CategoryRepository) (*[]categoriesRepository.Model, *customError.Error) {
	var result []categoriesRepository.Model
	//DBからデータを取得
	categories, err := r.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	// IDを取得するためのヘルパー関数
	var appendCategory func(category *categoriesRepository.Model)
	appendCategory = func(category *categoriesRepository.Model) {
		// まず、配列の先頭に自身を追加
		result = append([]categoriesRepository.Model{*category}, result...)
		if category.Parent != nil {
			appendCategory(FindCategoryByID(categories, *category.Parent))
		}
	}

	appendCategory(FindCategoryByID(categories, targetId))

	return &result, nil
}

// HasIdInTrail 指定されたidがtargetIdに至るまでの階層に存在するかチェック
func HasIdInTrail(ctx context.Context, r *categoriesRepository.CategoryRepository, id int, targetId int) (bool, *customError.Error) {
	// targetIdまでのパンくずリストを取得
	categoryTrail, err := GetCategoryTrail(ctx, targetId, r)
	if err != nil {
		return false, err
	}

	// 最後の要素を除いたパンくずリストを確認
	for _, category := range *categoryTrail {
		if category.ID == id {
			return true, nil
		}
	}

	return false, nil
}
