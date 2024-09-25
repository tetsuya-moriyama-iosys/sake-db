package bookmarkRepository

import (
	"backend/db"
	"backend/db/repository/agg"
	"backend/db/repository/userRepository"
	"errors"
	"fmt"
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
		UserID:           uid,
		BookmarkedUserId: targetId,
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
		agg.Where(UserID, uid),
		agg.LookUp(userRepository.CollectionName, BookmarkedUserId, userRepository.ID, "user_data"),
		agg.GetFirst("user_data"),
		bson.M{"$sort": bson.M{
			"_id": -1, // _idで降順ソート（新しい順）
		}},
		bson.M{"$addFields": bson.M{
			CreatedAT: bson.M{"$toDate": "$$ROOT._id"}, // _id からタイムスタンプを生成し、created_atフィールドに追加
		}},
		//projectで整形する
		bson.M{"$project": bson.M{
			UserID:    "$user_data." + userRepository.ID, // usersコレクションからのuser_name
			UserName:  "$user_data." + userRepository.NAME,
			CreatedAT: 1,
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

func (r *BookMarkRepository) GetRecommendLiquors(ctx context.Context, uid primitive.ObjectID, limitArg *int) ([]*Model, error) {
	//limit := 10
	//if limitArg != nil {
	//	limit = *limitArg
	//}

	fmt.Printf("uid: %+v\n", uid.Hex())
	pipeline := bson.A{
		// ドキュメントをフィルタリング
		agg.Where(ID, uid),
		//ユーザー情報を結合
		//agg.LookUp(userRepository.CollectionName, BookmarkedUserId, userRepository.ID, "user_data"),
		//agg.GetFirst("user_data"),
		////おすすめ情報を結合
		//agg.LookUp(liquorRepository.BoardCollectionName, BookmarkedUserId, liquorRepository.UserID, "recommend_data"),
		//bson.M{
		//	"$match": bson.M{
		//		"recommend_data." + liquorRepository.Rate: bson.M{"$gte": 4},
		//	},
		//},
		//agg.LookUp(liquorRepository.CollectionName, "recommend_data."+liquorRepository.LiquorID, liquorRepository.LiquorID, "liquor"),
		////agg.GetFirst("liquor"), // GetFirst を使用しないで重複排除
		//
		//// liquor_id でグループ化して重複を排除し、必要なフィールドを集計
		//bson.M{
		//	"$group": bson.M{
		//		"_id":         "$recommend_data." + liquorRepository.LiquorID,
		//		"liquor_data": bson.M{"$first": "$liquor"}, // 1:1の対応なのでfirstを使用
		//		"rate":        bson.M{"$first": "$recommend_data.rate"},
		//		"comment":     bson.M{"$first": "$recommend_data.comment"},
		//		"user_id":     bson.M{"$first": "$recommend_data.user_id"},
		//	},
		//},
		////  $sample でランダムに指定件数取得
		//bson.M{
		//	"$sample": bson.M{"size": limit},
		//},
	}

	// パイプライン実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するためのスライス
	var result []*Model

	// 結果を取得してコンソールに出力する
	for cursor.Next(ctx) {
		var doc bson.M
		//err := cursor.Decode(&doc)
		//if err != nil {
		//	return nil, err
		//}

		// 結果をコンソールに表示
		fmt.Printf("Document: %+v\n", doc)

		// 必要なら result に追加
		model := &Model{
			// doc から必要なフィールドをパースして model に追加する
		}
		result = append(result, model)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
