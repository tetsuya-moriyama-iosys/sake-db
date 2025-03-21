package errorRepository

import (
	"backend/db"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type ErrorsRepository struct {
	db         *db.DB
	collection *mongo.Collection //コレクションを先に取得して格納しておく
}

func New(db *db.DB) *ErrorsRepository {
	return &ErrorsRepository{
		db:         db,
		collection: db.Collection(CollectionName),
	}
}

func (r *ErrorsRepository) Write(ctx context.Context, errLog *Model) error {
	// MongoDBにデータを挿入
	_, err := r.collection.InsertOne(ctx, errLog)
	return err
}
