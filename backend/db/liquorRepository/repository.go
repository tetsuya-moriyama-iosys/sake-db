package liquorRepository

import (
	"backend/db"
	liquorModel "backend/graph/model/liquor"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LiquorsRepository struct {
	db         *db.DB
	collection *mongo.Collection //コレクションを先に取得して格納しておく
}

func NewLiquorsRepository(db *db.DB) LiquorsRepository {
	return LiquorsRepository{
		db:         db,
		collection: db.Collection(liquorModel.CollectionName),
	}
}

func (r *LiquorsRepository) InsertOne(ctx context.Context, liquor *liquorModel.Schema) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, liquor)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// InsertOneResultからIDを取得
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return id, nil
}
