package liquorRepository

import (
	"backend/db"
	"backend/graph/graphModel/liquor"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

func (r *LiquorsRepository) GetLiquorById(ctx context.Context, id string) (*liquorModel.Schema, error) {
	// idをObjectIDに変換
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("無効なID形式: %s", id)
	}

	// コレクションを取得
	var liquor liquorModel.Schema
	if err := r.collection.FindOne(ctx, bson.M{liquorModel.ID: objectID}).Decode(&liquor); err != nil {
		log.Println("デコードエラー:", err)
		return nil, err
	}

	return &liquor, nil
}

func (r *LiquorsRepository) GetRandomLiquors(ctx context.Context, limit int) ([]*liquorModel.Schema, error) {
	// $sampleパイプラインを使ってランダムに指定件数を取得
	cursor, err := r.collection.Aggregate(ctx, mongo.Pipeline{
		{{"$sample", bson.D{{"size", limit}}}},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var collections []*liquorModel.Schema
	if err = cursor.All(ctx, &collections); err != nil {
		return nil, err
	}

	return collections, nil
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
