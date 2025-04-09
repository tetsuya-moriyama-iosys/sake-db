package categoriesRepository

import (
	"backend/db"
	"backend/middlewares/customError"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (r *CategoryRepository) GetCategories(ctx context.Context) ([]*Model, *customError.Error) {
	//データを取得
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errFind(err)
	}
	defer cursor.Close(ctx)

	var categoryList []*Model
	for cursor.Next(ctx) {
		var category Model
		if err = cursor.Decode(&category); err != nil {
			return nil, errFindDecode(err)
		}
		categoryList = append(categoryList, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, errFindCursor(err)
	}

	return categoryList, nil
}

// GetCategoryByID IDからカテゴリを取得する
func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (*Model, *customError.Error) {
	var result Model
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errNotFound(err, id)
		}
		return nil, errFindById(err, id)
	}

	return &result, nil
}

func (r *CategoryRepository) InsertOne(ctx context.Context, category *Model) *customError.Error {
	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return errInsertOne(err, category)
	}

	// InsertOneResultからIDを取得
	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return errInsertOneGetId(err, category)
	}

	return nil
}

func (r *CategoryRepository) UpdateOne(ctx context.Context, category *Model) *customError.Error {
	// フィルタ：IDを用いてドキュメントを特定
	filter := bson.M{"id": category.ID}

	// 構造体を BSON にマッピング
	data, err := bson.Marshal(category)
	if err != nil {
		return errUpdateBsonMap(err, category)
	}

	// BSON を bson.M に変換
	var update bson.M
	if err := bson.Unmarshal(data, &update); err != nil {
		return errUpdateUnMarshal(err, category)
	}

	// 更新内容：$setオペレーターを使って指定したフィールドを更新
	updateBson := bson.M{"$set": update}

	// UpdateOneでドキュメントを更新
	result, err := r.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return errUpdateOne(err, category)
	}

	// UpdateOneは更新したドキュメントのIDを直接返さないため、元のIDを返す
	if result.MatchedCount == 0 {
		return errUpdateOneGetId(err, category)
	}

	return nil
}

// GetMaxID 最大のIDを取得する
func (r *CategoryRepository) GetMaxID(ctx context.Context) (int, *customError.Error) {
	// MongoDBのドキュメントのIDフィールドの最大値を取得するためのオプション
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	var result struct {
		ID int `bson:"id"`
	}

	err := r.collection.FindOne(ctx, bson.D{}, findOptions).Decode(&result)
	if err != nil {
		return 0, errGetMaxID(err)
	}

	return result.ID, nil
}
