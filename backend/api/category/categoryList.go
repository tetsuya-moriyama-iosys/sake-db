package category

import (
	"context"
	"log"
	"sort"
	"backend/db"

	"go.mongodb.org/mongo-driver/bson"
	
	"backend/graph/model"
)

//構造体の定義
type Category struct {
	ID     int `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	Parent *int `json:"parent" bson:"parent"`
	Children []*Category `json:"children,omitempty" bson:"-"`// 子カテゴリはDBに保存されないため、bsonタグは不要
}

func GetCategoryList() ([]*Category,error){
	//データベースに接続
	db.ConnectDB()

	//データを取得
	cursor, err := db.GetCollection("categories").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("データ取得エラー:", err)
        return nil,err
    }
	defer cursor.Close(context.TODO())

	categoryMap := make(map[int]*Category)
	var rootCategories []*Category

	for cursor.Next(context.TODO()) {
		var category Category
		if err := cursor.Decode(&category); err != nil {
			log.Fatal("デコードエラー:", err)
			return nil, err
		}
		categoryMap[category.ID] = &category
	}

	// 階層構造の構築
	for _, cat := range categoryMap {
		if cat.Parent == nil {
			rootCategories = append(rootCategories, cat)
		} else {
			parent := categoryMap[*cat.Parent]
			parent.Children = append(parent.Children, cat)
		}
	}

	// ルートカテゴリとその子カテゴリを再帰的にソート
	sortCategories(rootCategories)

	return rootCategories,nil
}

// 再帰的にカテゴリをID順にソートする関数
func sortCategories(categories []*Category) {
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].ID < categories[j].ID
	})
	for _, cat := range categories {
		if len(cat.Children) > 0 {
			sortCategories(cat.Children)
		}
	}
}


func ConvertToModelCategory(c *Category) *model.Category {
	mc := &model.Category{
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