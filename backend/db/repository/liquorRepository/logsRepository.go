package liquorRepository

import (
	"backend/db"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
