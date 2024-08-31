package helper

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv 環境変数の読み込み
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: loading .env file")
	}
}
