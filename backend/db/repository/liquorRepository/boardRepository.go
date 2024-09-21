package liquorRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *LiquorsRepository) BoardList(ctx context.Context, id primitive.ObjectID) ([]*BoardModelWithRelation, error) {
	// パイプラインを定義
	pipeline := bson.A{
		// 1. liquor_idに一致するドキュメントをフィルタリング
		bson.M{"$match": bson.M{"liquor_id": id}},

		// 2. usersコレクションとuser_idで結合してuser_nameを取得
		bson.M{"$lookup": bson.M{
			"from":         "users",     // 参照するコレクション
			"localField":   "user_id",   // boardのuser_idフィールド
			"foreignField": "_id",       // usersコレクションの_idフィールド
			"as":           "user_info", // 結果をuser_infoに格納
		}},

		// 3. 結果が配列なので、最初の要素に展開
		bson.M{"$unwind": bson.M{"path": "$user_info", "preserveNullAndEmptyArrays": true}},

		// 4. liquorsコレクションとliquor_idで結合してliquor_nameを取得
		bson.M{"$lookup": bson.M{
			"from":         "liquors",     // 参照するコレクション
			"localField":   "liquor_id",   // boardのliquor_idフィールド
			"foreignField": "_id",         // liquorsコレクションの_idフィールド
			"as":           "liquor_info", // 結果をliquor_infoに格納
		}},

		// 5. 結果が配列なので、最初の要素に展開
		bson.M{"$unwind": bson.M{"path": "$liquor_info", "preserveNullAndEmptyArrays": true}},

		// 6. 必要なフィールドだけをプロジェクト
		bson.M{"$project": bson.M{
			"_id":           1,
			"user_id":       1,
			"user_name":     "$user_info.name", // usersコレクションからのuser_name
			"liquor_id":     1,
			"liquor_name":   "$liquor_info.name", // liquorsコレクションからのliquor_name
			"category_id":   "$liquor_info.category_id",
			"category_name": "$liquor_info.category_name",
			"rate":          1,
			"text":          1,
			"updated_at":    1,
		}},
	}

	// パイプラインを実行
	cursor, err := r.boardCollection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var boards []*BoardModelWithRelation

	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &boards); err != nil {
		return nil, err
	}

	return boards, nil
}

// BoardListByUser ユーザーに紐づく掲示板投稿履歴を取得する。評価別および最近のものを取得
func (r *LiquorsRepository) BoardListByUser(ctx context.Context, uId primitive.ObjectID, limit int) (*BoardListResponse, error) {
	pipeline := bson.A{
		bson.M{"$match": bson.M{UserID: uId}}, // フィルタ
		bson.M{"$facet": bson.M{
			"groupedByRate": bson.A{
				bson.M{"$group": bson.M{
					"_id":   "$rate", // rateごとにグループ化
					"posts": bson.M{"$push": "$$ROOT"},
				}},
				bson.M{"$lookup": bson.M{
					"from":         CollectionName,    // 結合するコレクション
					"localField":   "posts.liquor_id", // groupされたドキュメントのliquor_id
					"foreignField": "_id",             // Liquorコレクションの_idフィールド
					"as":           "liquorDetails",   // 結合結果をliquorDetailsフィールドに格納
				}},
				// liquorDetailsをdocuments内のliquorフィールドとして追加
				bson.M{"$addFields": bson.M{
					"posts": bson.M{
						"$map": bson.M{
							"input": "$posts",
							"as":    "post",
							"in": bson.M{
								"_id":         "$$post._id",
								"category_id": "$$post.category_id",
								"liquor_id":   "$$post.liquor_id",
								"liquor_name": "$$post.liquor_name",
								"text":        "$$post.text",
								"updated_at":  "$$post.updated_at",
								// liquorDetailsの最初の要素をliquorフィールドとして埋め込む
								"liquor": bson.M{
									"$arrayElemAt": bson.A{"$liquorDetails", 0},
								},
							},
						},
					},
				}},
				// liquorDetailsを除外する
				bson.M{"$project": bson.M{
					"liquorDetails":           0, // liquorDetailsフィールドを除外
					"posts.category_id":       0,
					"posts.liquor_id":         0,
					"posts.liquor_name":       0,
					"posts.liquor.version_no": 0,
				}},
			},
			"recentPosts": bson.A{
				bson.M{"$sort": bson.M{ID: -1}}, // 降順にソート
				bson.M{"$limit": limit},         // 直近n件を取得
				// liquor_idでLiquorコレクションを$lookup
				bson.M{"$lookup": bson.M{
					"from":         CollectionName,  // 結合するコレクション
					"localField":   "liquor_id",     // recentDocuments内のliquor_id
					"foreignField": "_id",           // Liquorコレクションの_idフィールド
					"as":           "liquorDetails", // 結合結果をliquorDetailsフィールドに格納
				}},
				// liquorDetailsをrecentDocuments内のliquorフィールドとして追加
				bson.M{"$addFields": bson.M{
					"liquor": bson.M{
						"$arrayElemAt": bson.A{"$liquorDetails", 0}, // liquorDetailsの最初の要素をliquorフィールドとして埋め込む
					},
				}},
				// liquorDetailsを除外する
				bson.M{"$project": bson.M{
					"liquorDetails":     0, // liquorDetailsフィールドを除外
					"category_id":       0,
					"liquor_id":         0,
					"liquor_name":       0,
					"user_id":           0,
					"user_name":         0,
					"liquor.version_no": 0,
				}},
			},
		}},
	}

	// MongoDBの集計クエリを実行
	cursor, err := r.boardCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// クエリ結果を詰め替え
	var result *BoardListResponse
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// BoardGetByUserAndLiquor ユーザーIDとLiquorIDの組み合わせで、一意のモデルを取得する(編集用)
func (r *LiquorsRepository) BoardGetByUserAndLiquor(ctx context.Context, liquorId primitive.ObjectID, userId primitive.ObjectID) (*BoardModel, error) {
	// コレクションからフィルタに一致するドキュメントを取得
	var board *BoardModel
	if err := r.boardCollection.FindOne(ctx, bson.M{LiquorID: liquorId, UserID: userId}).Decode(&board); err != nil {
		return nil, err
	}

	return board, nil
}

func (r *LiquorsRepository) BoardInsert(ctx context.Context, board *BoardModel) error {
	// user_idが空かどうかを判定
	if board.UserId == nil {
		// user_idが空の場合はInsertOneを使用
		_, err := r.boardCollection.InsertOne(ctx, board)
		if err != nil {
			return err
		}
		return nil
	}
	// フィルタ：liquorID,userIDが既に存在するか確認
	filter := bson.M{
		UserID:   board.UserId,
		LiquorID: board.LiquorID,
	}

	// 更新データ：board内のフィールドをそのまま更新
	update := bson.M{
		"$set": board,
	}

	// upsertオプション：ドキュメントが存在しない場合は新規挿入
	opts := options.Update().SetUpsert(true)

	// MongoDBにデータを更新または挿入（upsert）
	_, err := r.boardCollection.UpdateOne(ctx, filter, update, opts)
	//_, err := r.boardCollection.InsertOne(ctx, board)
	if err != nil {
		return err
	}

	return nil
}
