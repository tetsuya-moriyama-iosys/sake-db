package main

import (
	"backend/db/indexes"
	"backend/di"
	"backend/util/helper"
	"log"
)

func main() {
	helper.LoadEnv() //.envファイルを読み込み可能にする
	server, err := di.InitializeHandler()
	if err != nil {
		log.Fatal("Failed to initialize server:", err)
	}
	// インデックス作成処理を呼び出す
	err = indexes.AddIndexes()
	if err != nil {
		log.Fatal("Failed to initialize indexes:", err)
	}

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(server.Run(":8080"))
}
