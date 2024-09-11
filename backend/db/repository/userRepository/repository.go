package userRepository

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UsersRepository struct {
	db         *db.DB
	collection *mongo.Collection //コレクションを先に取得して格納しておく
}

func NewUsersRepository(db *db.DB) UsersRepository {
	return UsersRepository{
		db:         db,
		collection: db.Collection(CollectionName),
	}
}

func (r *UsersRepository) Register(ctx context.Context, user *Model) (*Model, error) {
	// MongoDBにデータを挿入
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// 挿入されたIDをuserのIDにセット
	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*Model, error) {
	// ドキュメントを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{EMAIL: email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) GetById(ctx context.Context, id string) (*Model, error) {
	// idをObjectIDに変換
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("無効なID形式: %s", id)
	}

	// コレクションを取得
	var liquor Model
	if err := r.collection.FindOne(ctx, bson.M{ID: objectID}).Decode(&liquor); err != nil {
		log.Println("デコードエラー:", err)
		return nil, err
	}

	return &liquor, nil
}
