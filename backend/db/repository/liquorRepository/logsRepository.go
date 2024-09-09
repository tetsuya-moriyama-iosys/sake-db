package liquorRepository

import (
	"backend/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *LiquorsRepository) GetLogsById(ctx context.Context, id string) ([]*Model, error) {
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

func (r *LiquorsRepository) GetLogsByVersionNo(ctx context.Context, id string, versionNo int) (*Model, error) {
	// カテゴリIDがidのコレクションを降順で取得
	var model *Model
	err := r.logsCollection.FindOne(ctx, bson.M{"id": id, VersionNo: versionNo}).Decode(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *LiquorsRepository) InsertOneToLog(ctx context.Context, liquor *Model) error {
	// 既存の _id を id フィールドに移動
	liquorID := liquor.ID               // 現在の _id を保存
	liquor.ID = primitive.NewObjectID() // _id に新しい ObjectID を割り当て

	data, err := db.StructToBsonM(liquor)
	if err != nil {
		return err
	}

	data[LogID] = liquorID.Hex()

	// ログコレクションに挿入
	_, err = r.logsCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
