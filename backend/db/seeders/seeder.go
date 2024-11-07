package main

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/flavorMapRepository"
	"backend/util/helper"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	helper.LoadEnv() // .envファイルを読み込み可能にする

	// MongoDBのクライアントを作成
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// シードデータの設定
	dbName := os.Getenv("MAIN_DB_NAME")
	ctx := context.Background()

	// Categoryモデルのシード
	var categoryModel []map[string]interface{}
	err = seedData(ctx, client, dbName, categoriesRepository.CollectionName, "./db/seeders/categories.json", &categoryModel, "id")
	if err != nil {
		log.Fatal(err)
	}

	// 他のモデルのシード
	var flavorMapMstModel []map[string]interface{}
	err = seedData(ctx, client, dbName, flavorMapRepository.FlavorMapMasterCollectionName, "./db/seeders/flavorMaps.json", &flavorMapMstModel, "category_id")
	if err != nil {
		log.Fatal(err)
	}
}

// SeedData シード処理の共通関数
func seedData(ctx context.Context, client *mongo.Client, dbName, collectionName, filePath string, model interface{}, idField string) error {
	// データベースとコレクションの参照を取得
	collection := client.Database(dbName).Collection(collectionName)

	// JSONファイルからデータを読み込む
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// デコード
	if err := json.NewDecoder(file).Decode(model); err != nil {
		return fmt.Errorf("failed to decode JSON from file %s: %w", filePath, err)
	}

	// デコードしたデータを配列として取得
	items := *model.(*[]map[string]interface{}) // ポインタをデリファレンスして、[]map[string]interface{}型に変換

	// 各ドキュメントをアップサート
	for _, item := range items {
		// _idフィールドが存在し、かつ$oid形式であればObjectIDに変換
		if id, ok := item["_id"].(map[string]interface{}); ok {
			if oidStr, ok := id["$oid"].(string); ok {
				objectID, err := primitive.ObjectIDFromHex(oidStr)
				if err != nil {
					return fmt.Errorf("invalid ObjectID format in _id field: %w", err)
				}
				item["_id"] = objectID
			}
		}

		// アップサート用のfilterとupdateを作成
		filter := bson.M{idField: item[idField]}

		// _idフィールドを除外したupdate用のコピーを作成
		itemWithoutID := make(map[string]interface{})
		for k, v := range item {
			if k != "_id" {
				itemWithoutID[k] = v
			}
		}

		// 新規挿入時のみ_idを設定
		update := bson.M{
			"$set":         itemWithoutID,
			"$setOnInsert": bson.M{"_id": item["_id"]},
		}

		// Upsertオプションを設定
		options := options.Update().SetUpsert(true)

		// アップサートを実行
		result, err := collection.UpdateOne(ctx, filter, update, options)
		if err != nil {
			return fmt.Errorf("failed to upsert document: %w", err)
		}

		if result.MatchedCount > 0 {
			fmt.Printf("%s Updated document with %s: %v\n", collectionName, idField, item[idField])
		} else {
			fmt.Printf("%s Inserted new document with %s: %v\n", collectionName, idField, item[idField])
		}
	}

	return nil
}
