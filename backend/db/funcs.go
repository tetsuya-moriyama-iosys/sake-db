package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *Base) Upsert(ctx context.Context, filter bson.M, model interface{}) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)

	// update フィールドを自動的に生成
	update := bson.M{"$set": model}

	result, err := b.Collection.UpdateOne(ctx, filter, update, opts)
	return result, err
}
