package liquorRepository

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type LiquorsRepository struct {
	db             *db.DB
	collection     *mongo.Collection //コレクションを先に取得して格納しておく
	logsCollection *mongo.Collection //コレクションを先に取得して格納しておく
}

func NewLiquorsRepository(db *db.DB) LiquorsRepository {
	return LiquorsRepository{
		db:             db,
		collection:     db.Collection(CollectionName),
		logsCollection: db.Collection(LogsCollectionName),
	}
}

func (r *LiquorsRepository) GetLiquorById(ctx context.Context, id string) (*Model, error) {
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

func (r *LiquorsRepository) GetRandomLiquors(ctx context.Context, limit int) ([]*Model, error) {
	// $sampleパイプラインを使ってランダムに指定件数を取得
	cursor, err := r.collection.Aggregate(ctx, mongo.Pipeline{
		{{"$sample", bson.D{{"size", limit}}}},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var collections []*Model
	if err = cursor.All(ctx, &collections); err != nil {
		return nil, err
	}

	return collections, nil
}

func (r *LiquorsRepository) GetLiquorsFromCategoryIds(ctx context.Context, ids []int) ([]*Model, error) {
	// クエリフィルターを作成。カテゴリIDがidsのいずれかに一致するリカーを取得
	filter := bson.M{"category_id": bson.M{"$in": ids}}

	// コレクションからフィルタに一致するドキュメントを取得
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var liquors []*Model

	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &liquors); err != nil {
		return nil, err
	}

	return liquors, nil
}

func (r *LiquorsRepository) InsertOne(ctx context.Context, liquor *Model) (primitive.ObjectID, error) {
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

func (r *LiquorsRepository) UpdateOne(ctx context.Context, liquor *Model) (primitive.ObjectID, error) {
	// フィルタ：IDを用いてドキュメントを特定
	filter := bson.M{"_id": liquor.ID}

	// 構造体を BSON にマッピング
	data, err := bson.Marshal(liquor)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// BSON を bson.M に変換
	var update bson.M
	if err := bson.Unmarshal(data, &update); err != nil {
		return primitive.NilObjectID, err
	}

	// 更新内容：$setオペレーターを使って指定したフィールドを更新
	updateBson := bson.M{"$set": update}

	// UpdateOneでドキュメントを更新
	result, err := r.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// UpdateOneは更新したドキュメントのIDを直接返さないため、元のIDを返す
	if result.MatchedCount == 0 {
		return primitive.NilObjectID, fmt.Errorf("no document matched the provided ID")
	}

	return liquor.ID, nil
}