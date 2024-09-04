package categoriesRepository

import (
	"backend/db"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

type CategoryRepository struct {
	db             *db.DB
	collection     *mongo.Collection
	logsCollection *mongo.Collection
}

func NewCategoryRepository(db *db.DB) CategoryRepository {
	return CategoryRepository{
		db:             db,
		collection:     db.Collection(CollectionName),
		logsCollection: db.Collection(LogsCollectionName),
	}
}

// GetCategories カテゴリの一覧を取得する
func (r *CategoryRepository) GetCategories(ctx context.Context) ([]*Model, error) {
	//データを取得
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("データ取得エラー:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var categoryList []*Model
	for cursor.Next(ctx) {
		var category Model
		if err := cursor.Decode(&category); err != nil {
			log.Fatal("デコードエラー:", err)
			return nil, err
		}
		categoryList = append(categoryList, &category)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("カーソルエラー:", err)
		return nil, err
	}

	return categoryList, nil
}

// GetCategoryByID IDからカテゴリを取得する
func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (*Model, error) {
	var result Model
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &result, nil
}

func (r *CategoryRepository) InsertOne(ctx context.Context, category *Model) error {
	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	// InsertOneResultからIDを取得
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return err
	}

	return nil
}

func (r *CategoryRepository) UpdateOne(ctx context.Context, liquor *Model) error {
	// フィルタ：IDを用いてドキュメントを特定
	filter := bson.M{"id": liquor.ID}

	// 構造体を BSON にマッピング
	data, err := bson.Marshal(liquor)
	if err != nil {
		return err
	}

	// BSON を bson.M に変換
	var update bson.M
	if err := bson.Unmarshal(data, &update); err != nil {
		return err
	}

	// 更新内容：$setオペレーターを使って指定したフィールドを更新
	updateBson := bson.M{"$set": update}

	// UpdateOneでドキュメントを更新
	result, err := r.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return err
	}

	// UpdateOneは更新したドキュメントのIDを直接返さないため、元のIDを返す
	if result.MatchedCount == 0 {
		return errors.New("no document matched the provided ID")
	}

	return nil
}

func (r *CategoryRepository) InsertOneToLog(ctx context.Context, category *Model) error {
	data, err := db.StructToBsonM(category)
	if err != nil {
		return err
	}

	//data["_id"] = primitive.NewObjectID()//←必要なかったら消す

	// ログコレクションに挿入
	_, err = r.logsCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

// GetMaxID 最大のIDを取得する
func (r *CategoryRepository) GetMaxID(ctx context.Context) (int, error) {
	// MongoDBのドキュメントのIDフィールドの最大値を取得するためのオプション
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	var result struct {
		ID int `bson:"id"`
	}

	err := r.collection.FindOne(ctx, bson.D{}, findOptions).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// コレクションにドキュメントがない場合、最大IDを0とする
			return 0, err
		}
		return 0, err
	}

	return result.ID, nil
}
