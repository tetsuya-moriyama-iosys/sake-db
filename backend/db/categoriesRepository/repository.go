package categoriesRepository

import (
	"backend/db"
	categoryModel "backend/graph/model/category"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
)

type CategoryRepository struct {
	db         *db.DB
	collection *mongo.Collection
}

// Category 構造体の定義
type Category struct {
	ID       int         `json:"id" bson:"id"`
	Name     string      `json:"name" bson:"name"`
	Parent   *int        `json:"parent" bson:"parent"`
	Children []*Category `json:"children,omitempty" bson:"-"` // 子カテゴリはDBに保存されないため、bsonタグは不要
}

func NewCategoryRepository(db *db.DB) CategoryRepository {
	return CategoryRepository{
		db:         db,
		collection: db.Collection(categoryModel.CollectionName),
	}
}

// GetCategories カテゴリの一覧を取得する
func (r *CategoryRepository) GetCategories(ctx context.Context) ([]*Category, error) {
	//データを取得
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("データ取得エラー:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var categoryMap []*Category
	for cursor.Next(context.TODO()) {
		var category Category
		if err := cursor.Decode(&category); err != nil {
			log.Fatal("デコードエラー:", err)
			return nil, err
		}
		categoryMap[category.ID] = &category
	}

	return categoryMap, nil
}

// GetCategoryNameByID IDからカテゴリ名を取得する
func (r *CategoryRepository) GetCategoryNameByID(ctx context.Context, id int) (string, error) {
	log.Println("category_id is:", id)
	var result bson.M
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	if err != nil {
		log.Println("not found")
		if err == mongo.ErrNoDocuments {
			return "", errors.New("category not found")
		}
		return "", err
	}
	log.Println("found:", result)

	name, ok := result[categoryModel.Name].(string)
	if !ok {
		return "", errors.New("name field is not found or not a string")
	}

	return name, nil
}
