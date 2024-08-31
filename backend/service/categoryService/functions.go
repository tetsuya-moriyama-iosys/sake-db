package categoryService

import (
	"backend/db/categoriesRepository"
	"backend/graph/graphModel"
	"sort"
)

// 再帰的にカテゴリをID順にソートする関数
func sortCategories(categories []*categoriesRepository.Category) {
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].ID < categories[j].ID
	})
	for _, cat := range categories {
		if len(cat.Children) > 0 {
			sortCategories(cat.Children)
		}
	}
}

func ConvertToModelCategory(c *categoriesRepository.Category) *graphModel.Category {
	mc := &graphModel.Category{
		ID:   c.ID,
		Name: c.Name,
	}
	if c.Parent != nil {
		mc.Parent = c.Parent
	}
	for _, child := range c.Children {
		mc.Children = append(mc.Children, ConvertToModelCategory(child))
	}
	return mc
}
