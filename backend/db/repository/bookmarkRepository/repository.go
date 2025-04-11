package bookmarkRepository

import (
	"backend/db"
	"backend/db/repository/agg"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"backend/middlewares/customError"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *BookMarkRepository) Find(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) (*Model, *customError.Error) {
	// クエリを実行し、ドキュメントが存在するか確認
	var result *Model
	err := r.collection.FindOne(ctx, filter(uid, targetId)).Decode(&result) //取得しデコードする
	if err != nil {
		return nil, errFindOne(err, uid, targetId)
	}
	return result, nil
}

func (r *BookMarkRepository) List(ctx context.Context, uid primitive.ObjectID) ([]*BookMarkListUser, *customError.Error) {
	// パイプラインを定義
	pipeline := generatePipeline(uid, UserId, BookmarkedId)

	// パイプラインを実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errListAggregate(err, uid)
	}
	defer cursor.Close(ctx)
	// 結果を格納するスライス
	var bList []*BookMarkListUser
	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &bList); err != nil {
		return nil, errListDecode(err, uid)
	}
	return bList, nil
}

// BookmarkedList そのユーザーをブックマークしている人のリスト
func (r *BookMarkRepository) BookmarkedList(ctx context.Context, uid primitive.ObjectID) ([]*BookMarkListUser, *customError.Error) {
	// パイプラインを定義(Listとは逆に 、被ブックマークIDで絞り込む)
	pipeline := generatePipeline(uid, BookmarkedId, UserId)
	// パイプラインを実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errBookmarkedListAggregate(err, uid)
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var bList []*BookMarkListUser
	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &bList); err != nil {
		return nil, errBookmarkedListDecode(err, uid)
	}
	return bList, nil
}

func (r *BookMarkRepository) Add(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) *customError.Error {
	_, findErr := r.Find(ctx, uid, targetId)
	if findErr == nil {
		//見つかった場合は重複するのでエラー
		return errDuplicated(findErr, uid, targetId)
	}

	if !errors.Is(findErr, mongo.ErrNoDocuments) {
		// ドキュメントが存在しない以外のエラーは普通にエラーなので返す
		return errOnAddFind(findErr, uid, targetId)
	}

	//レコードを挿入する
	_, err := r.collection.InsertOne(ctx, &Model{
		UserId:           uid,
		BookmarkedUserId: targetId,
	})
	if err != nil {
		return errOnAdd(findErr, uid, targetId)
	}
	return nil
}

func (r *BookMarkRepository) Remove(ctx context.Context, uid primitive.ObjectID, targetId primitive.ObjectID) *customError.Error {
	//レコードを削除する
	result, err := r.collection.DeleteOne(ctx, filter(uid, targetId))
	if err != nil {
		return errDeleteOne(err, uid, targetId)
	}

	if result.DeletedCount == 0 {
		return errOnDelete(err, uid, targetId)
	}
	return nil
}

func (r *BookMarkRepository) GetRecommendLiquors(ctx context.Context, uid primitive.ObjectID, limitArg *int) (*RecommendList, *customError.Error) {
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
		agg.LookUp(userRepository.CollectionName, "user_id", userRepository.Id, "user_info"),
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
				"user_info._id":          1,
				"user_info.name":         1,
				"user_info.image_base64": 1,
			},
		},

		bson.M{
			"$sample": bson.M{"size": limit},
		},
	}

	// パイプライン実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errFindRecommend(err, uid)
	}
	defer cursor.Close(ctx)

	// 結果を格納するためのスライス
	var result RecommendList

	// 結果を取得してコンソールに出力する
	for cursor.Next(ctx) {
		var doc *Recommend
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, errRecommendDecode(err, uid)
		}

		// 結果をコンソールに表示
		//helper.D(doc)

		result = append(result, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, errRecommend(err, uid)
	}

	return &result, nil
}
