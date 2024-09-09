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
}

func AddIndexes() error {
	helper.LoadEnv()
	client, err := db.NewMongoClient()

	// 各コレクションに対してインデックスを作成
	for _, indexDef := range IndexDefinitions {
		collection := getDatabaseAndCollection(client, indexDef.CollectionName)
		err = createIndexForCollection(context.Background(), collection, indexDef.IndexKeys)
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
func createIndexForCollection(ctx context.Context, collection *mongo.Collection, indexKeys bson.D) error {
	indexModel := mongo.IndexModel{
		Keys:    indexKeys,
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}
