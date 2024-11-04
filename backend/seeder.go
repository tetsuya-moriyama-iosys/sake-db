package main

import (
	"backend/db/repository/categoriesRepository"
	"backend/util/helper"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	helper.LoadEnv() //.envファイルを読み込み可能にする
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

	// デコードした構造体スライスをそのまま MongoDB に挿入する場合、Go では []interface{} に型変換が必要
	documents := make([]interface{}, len(items))
	for i, item := range items {
		documents[i] = item
	}

	// 挿入
	result, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted documents: %v\n", result.InsertedIDs)
}
