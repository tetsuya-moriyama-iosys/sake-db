package liquorRepository

import (
	"backend/db"
	"backend/middlewares/customError"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *LiquorsRepository) GetLogsById(ctx context.Context, id primitive.ObjectID) ([]*Model, *customError.Error) {
	// liquorIDがidのコレクションを降順で取得
	cursor, err := r.logsCollection.Find(ctx, bson.M{LiquorID: id}, options.Find().SetSort(bson.D{{VersionNo, -1}}))
	if err != nil {
		return nil, errGetLogsById(err, id)
	}

	var result []*Model
	if err = cursor.All(ctx, &result); err != nil {
		return nil, errCursor(err, id)
	}
	defer cursor.Close(ctx)

	return result, nil
}

func (r *LiquorsRepository) GetLogsByVersionNo(ctx context.Context, id string, versionNo int) (*Model, *customError.Error) {
	// バージョンnoを指定して取得
	var model *Model
	err := r.logsCollection.FindOne(ctx, bson.M{LiquorID: id, VersionNo: versionNo}).Decode(&model)
	if err != nil {
		return nil, errGetLogsByVer(err, id, versionNo)
	}

	return model, nil
}

func (r *LiquorsRepository) InsertOneToLog(ctx context.Context, oldLiquor *Model) *customError.Error {
	newLiquor := *oldLiquor                //旧値を値コピー
	newLiquor.ID = primitive.NewObjectID() // _id に新しい ObjectID を割り当て

	//継承ができないので、interface型でlog用モデルを定義し直す(Liquorモデルにliquor_id(元となったliquor_id)を追加するだけなのでこの方がラク)
	data, err := db.StructToBsonM(newLiquor)
	if err != nil {
		return errToBsonForInsert(err)
	}

	data[LiquorID] = oldLiquor.ID

	// ログコレクションに挿入
	_, err = r.logsCollection.InsertOne(ctx, data)
	if err != nil {
		return errInsertLogOne(err, oldLiquor)
	}

	return nil
}
