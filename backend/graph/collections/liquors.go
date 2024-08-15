package collections

import (
	"backend/db"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Liquors() *mongo.Collection {
	db.ConnectDB()
	if db.Client == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return db.GetCollection("liquors")
}
