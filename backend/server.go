package main

import (
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

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(server.Run(":8080"))
}
