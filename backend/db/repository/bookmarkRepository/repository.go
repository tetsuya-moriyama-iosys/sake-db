package bookmarkRepository

import (
	"backend/db"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type BookMarkRepository struct {
	db         *db.DB
	collection *mongo.Collection
}

func NewBookMarkRepository(db *db.DB) BookMarkRepository {
	return BookMarkRepository{
		db:         db,
		collection: db.Collection(CollectionName),
	}
}

func (r *BookMarkRepository) Add(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) error {
	// クエリ条件
	filter := bson.M{
		USER_ID:            uid,
		BOOKMARKED_USER_ID: targetId,
	}
	// クエリを実行し、ドキュメントが存在するか確認
	var result bson.M
	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		//見つかった場合は重複するのでエラー
		return errors.New("すでにブックマークされています")
	}

	if err != mongo.ErrNoDocuments {
		// ドキュメントが存在しない以外のエラーは普通にエラーなので返す
		return err
	}

	//レコードを挿入する
	_, err = r.collection.InsertOne(ctx, &Model{
		UserId:           uid,
		BookmarkedUserId: targetId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *BookMarkRepository) Remove(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) error {
	// クエリ条件
	filter := bson.M{
		USER_ID:            uid,
		BOOKMARKED_USER_ID: targetId,
	}

	//レコードを削除する
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("ブックマークが存在しませんでした")
	}
	return nil
}
