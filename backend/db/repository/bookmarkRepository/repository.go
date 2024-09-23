package bookmarkRepository

import (
	"backend/db"
	"backend/db/repository/agg"
	"backend/db/repository/userRepository"
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

// 検索条件フィルター生成を共通化
func filter(uid primitive.ObjectID, targetId primitive.ObjectID) bson.M {
	return bson.M{
		USER_ID:            uid,
		BOOKMARKED_USER_ID: targetId,
	}
}

func (r *BookMarkRepository) Find(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) (*Model, error) {
	// クエリを実行し、ドキュメントが存在するか確認
	var result *Model
	err := r.collection.FindOne(ctx, filter(uid, targetId)).Decode(&result) //取得しデコードする
	//エラーごと返す(エラーならresultがnilのはず)
	return result, err
}

func (r *BookMarkRepository) List(ctx context.Context, uid primitive.ObjectID) ([]*BookMarkListUser, error) {
	// パイプラインを定義
	pipeline := bson.A{
		// ドキュメントをフィルタリング
		agg.Where(USER_ID, uid),
		agg.LookUp(userRepository.CollectionName, USER_ID, userRepository.ID, "user_data"),
		agg.GetFirst("user_data"),
		//projectで整形する
		bson.M{"$project": bson.M{
			USER_ID:   "$user_data." + userRepository.ID, // usersコレクションからのuser_name
			USER_NAME: "$user_data." + userRepository.NAME,
		}},
	}
	// パイプラインを実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	// 結果を格納するスライス
	var bList []*BookMarkListUser
	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &bList); err != nil {
		return nil, err
	}
	return bList, nil
}

func (r *BookMarkRepository) Add(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) error {
	_, err := r.Find(ctx, uid, targetId)
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
	//レコードを削除する
	result, err := r.collection.DeleteOne(ctx, filter(uid, targetId))
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("ブックマークが存在しませんでした")
	}
	return nil
}
