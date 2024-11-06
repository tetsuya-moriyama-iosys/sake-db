package main

import (
	"backend/db/repository/categoriesRepository"
	"backend/util/helper"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

	// データベースとコレクションの参照を取得
	collection := client.Database(os.Getenv("MAIN_DB_NAME")).Collection(categoriesRepository.CollectionName)

	// JSONファイルからデータを読み込む
	categoriesFile, err := os.Open("./db/seeders/categories.json")
	if err != nil {
		log.Fatal(err)
	}
	defer categoriesFile.Close()

	// デコードして挿入
	var items []categoriesRepository.Model
	if err := json.NewDecoder(categoriesFile).Decode(&items); err != nil {
		log.Fatal(err)
	}

	// デコードした各ドキュメントをアップサート
	for _, item := range items {
		filter := bson.M{"id": item.ID} // `id`が一致するかどうかでフィルター
		update := bson.M{"$set": item}  // ドキュメント全体を上書き

		// Upsertオプションを設定
		options := options.Update().SetUpsert(true)

		// アップサートを実行
		result, err := collection.UpdateOne(context.Background(), filter, update, options)
		if err != nil {
			log.Fatal(err)
		}

		if result.MatchedCount > 0 {
			fmt.Printf("Updated document with id: %v\n", item.ID)
		} else {
			fmt.Printf("Inserted new document with id: %v\n", item.ID)
		}
	}
}
