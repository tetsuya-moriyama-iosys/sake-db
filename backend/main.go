package main

import (
	"backend/db/indexes"
	"backend/di"
	"backend/middlewares"
	"backend/util/helper"
	"log"
)

func main() {
	helper.LoadEnv() //.envファイルを読み込み可能にする
	server, err := di.InitializeHandler()
	if err != nil {
		log.Fatal("Failed to initialize server:", err)
	}

	// ミドルウェアを登録
	server.Use(middlewares.AttachGinContextToContext())

	// インデックス作成処理を呼び出す
	err = indexes.AddIndexes()
	if err != nil {
		log.Fatal("Failed to initialize indexes:", err)
	}

	log.Println("connect to http://localhost:8080/query for GraphQL playground")
	log.Println(server.Run(":8080"))
}
