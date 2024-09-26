package bookmarkRepository

import (
	"backend/db"
	"backend/db/repository/agg"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"backend/util/helper"
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
		agg.GetFirst("user_data", false),
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
	limit := 10
	if limitArg != nil {
		limit = *limitArg
	}

	fmt.Printf("uid: %+v\n", uid.Hex())
	pipeline := bson.A{
		// ドキュメントをフィルタリング
		agg.Where(UserID, uid),
		//ブックマーク済ユーザーが投稿したおすすめ情報を結合
		agg.LookUp(liquorRepository.BoardCollectionName, BookmarkedUserId, liquorRepository.UserID, "recommend_data"),
		bson.M{
			"$addFields": bson.M{
				"recommend_data": bson.M{
					"$filter": bson.M{
						"input": "$recommend_data",
						"as":    "item",
						"cond":  bson.M{"$gte": []interface{}{"$$item.rate", 4}}, // rateが4以上のものを残す
					},
				},
			},
		},

		//おすすめ情報を展開して一次元にする
		bson.M{"$unwind": bson.M{"path": "$recommend_data"}},

		//liquor_id単位でグループ化
		bson.M{
			"$group": bson.M{
				"_id": "$recommend_data.liquor_id", // liquor_id ごとにグループ化
				"items": bson.M{
					"$push": "$$ROOT", // 各グループに全アイテムを配列で格納
				},
			},
		},
		// 各グループの items 配列からランダムに 1 件を取得
		bson.M{
			"$addFields": bson.M{
				"random_item": bson.M{
					"$arrayElemAt": bson.A{
						"$items", bson.M{
							"$floor": bson.M{
								"$multiply": bson.A{
									bson.M{"$rand": bson.M{}}, bson.M{"$size": "$items"},
								},
							},
						},
					},
				},
			},
		},

		//トップレベルのデータを整形する
		bson.M{
			"$addFields": bson.M{
				"liquor_id":  "$random_item.recommend_data.liquor_id",
				"rate":       "$random_item.recommend_data.rate",
				"comment":    "$random_item.recommend_data.text",
				"user_id":    "$random_item.recommend_data.user_id",
				"updated_at": "$random_item.recommend_data.updated_at",
			},
		},

		//ユーザー情報・お酒情報を結合
		agg.LookUp(userRepository.CollectionName, "user_id", userRepository.ID, "user_info"),
		agg.GetFirst("user_info", false),
		agg.LookUp(liquorRepository.CollectionName, "liquor_id", liquorRepository.ID, "liquor"),
		agg.GetFirst("liquor", false),

		// 必要なフィールドのみを取得する
		bson.M{
			"$project": bson.M{
				"_id":        0,
				"rate":       1,
				"comment":    1,
				"updated_at": 1,
				//liquor内の必要なフィールドを定義
				"liquor._id":           1,
				"liquor.name":          1,
				"liquor.category_id":   1,
				"liquor.category_name": 1,
				"liquor.image_base64":  1,
				"liquor.description":   1,
				//userテーブルについても同様
				"user_info._id":  1,
				"user_info.name": 1,
			},
		},

		bson.M{
			"$sample": bson.M{"size": limit},
		},
	}
	fmt.Printf("Pipeline: %+v\n", pipeline)

	// パイプライン実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するためのスライス
	var result []interface{} //一旦any型に

	// 結果を取得してコンソールに出力する
	for cursor.Next(ctx) {
		var doc bson.M
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}

		// 結果をコンソールに表示
		helper.D(doc)

		// 必要なら result に追加
		//model := &Model{
		//	// doc から必要なフィールドをパースして model に追加する
		//}
		result = append(result, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
