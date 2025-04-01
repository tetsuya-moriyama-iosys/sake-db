package categoriesRepository

import (
	"backend/db"
	"backend/middlewares/customError"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *CategoryRepository) GetLogsById(ctx context.Context, id int) ([]*Model, *customError.Error) {
	// カテゴリIDがidのコレクションを降順で取得
	cursor, err := r.logsCollection.Find(ctx, bson.M{ID: id}, options.Find().SetSort(bson.D{{VersionNo, -1}}))
	if err != nil {
		return nil, err
	}

	var result []*Model
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	return result, nil
}

func (r *CategoryRepository) GetLogsByVersionNo(ctx context.Context, id int, versionNo int) (*Model, *customError.Error) {
	// カテゴリIDがidのコレクションを降順で取得
	var model *Model
	err := r.logsCollection.FindOne(ctx, bson.M{"id": id, VersionNo: versionNo}).Decode(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *CategoryRepository) InsertOneToLog(ctx context.Context, category *Model) *customError.Error {
	data, err := db.StructToBsonM(category)
	if err != nil {
		return err
	}

	// ログコレクションに挿入
	_, err = r.logsCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
