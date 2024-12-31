package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Collection(name string) *mongo.Collection
}

// DB 接続情報を格納する構造体
type DB struct {
	Client *mongo.Client
	DBName string
}

type Base struct {
	Db         *DB
	Collection *mongo.Collection
}

func NewMongoClient() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewDB(client *mongo.Client) *DB {
	dbName := os.Getenv("MAIN_DB_NAME")
	return &DB{
		Client: client,
		DBName: dbName,
	}
}

func ProvideMongoDatabase(db *DB) *mongo.Database {
	return db.Client.Database(db.DBName)
}

func (db *DB) Collection(name string) *mongo.Collection {
	return db.Client.Database(db.DBName).Collection(name)
}
