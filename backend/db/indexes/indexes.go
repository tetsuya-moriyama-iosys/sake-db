package indexes

import (
	"backend/db"
	"backend/util/helper"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// IndexDefinition コレクション名とそのインデックスの定義を持つ構造体
type IndexDefinition struct {
	CollectionName string
	IndexKeys      bson.D
	IsNonUnique    bool   //未指定がfalseなので、NonUniqueとしている(基本はユニーク制約をつける想定)
	PartialFilter  bson.D // Optional: nullの場合ユニーク制約を外すためのフィルター
}

func AddIndexes() error {
	helper.LoadEnv()

	// 各コレクションに対してインデックスを作成
	for _, indexDef := range IndexDefinitions {
		err := createIndexForCollection(context.Background(), indexDef)
		if err != nil {
			return err
		}
	}

	return nil
}

// getDatabaseAndCollectionはデータベースとコレクションを取得します
func getDatabaseAndCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	db := client.Database(os.Getenv("MAIN_DB_NAME"))
	return db.Collection(collectionName)
}

// createIndexForCollectionは指定されたコレクションにインデックスを作成します
func createIndexForCollection(ctx context.Context, indexData IndexDefinition) error {
	client, err := db.NewMongoClient()
	collection := getDatabaseAndCollection(client, indexData.CollectionName)

	// インデックス作成のオプション設定
	indexOptions := options.Index().SetUnique(!indexData.IsNonUnique)

	// PartialFilterが設定されている場合、partialFilterExpressionを追加
	if len(indexData.PartialFilter) > 0 {
		indexOptions.SetPartialFilterExpression(indexData.PartialFilter)
	}

	indexModel := mongo.IndexModel{
		Keys:    indexData.IndexKeys,
		Options: indexOptions,
	}

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	return err
}
