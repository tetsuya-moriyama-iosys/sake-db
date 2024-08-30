package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// データベース名
var dbName = "helloworld"

func ConnectDB() {
	// 認証情報を含むMongoDB接続URI
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/helloworld?authSource=admin")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	Client = client
	log.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		ConnectDB()
	}
	return Client.Database(dbName).Collection(collectionName)
}
