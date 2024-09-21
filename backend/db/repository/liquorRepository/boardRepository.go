package liquorRepository

import (
	"backend/graph/graphModel"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	BoardCollectionName = "liquors_boards"
	Rate                = "rate"
	LiquorID            = "liquor_id"
	LiquorName          = "liquor_name"
	UserID              = "user_id"
	UserName            = "user_name"
)

type BoardModel struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	CategoryID int                 `bson:"category_id"`
	LiquorID   primitive.ObjectID  `bson:"liquor_id"`
	LiquorName string              `bson:"liquor_name"`
	UserId     *primitive.ObjectID `bson:"user_id"`
	UserName   *string             `bson:"user_name"`
	Text       string              `bson:"text"`
	Rate       *int                `bson:"rate"`
	UpdatedAt  time.Time           `bson:"updated_at"`
}

// Document 各投稿の詳細
type Document struct {
	Text      string       `bson:"text"`       // 投稿内容
	UpdatedAt time.Time    `bson:"updated_at"` // 更新日時
	Liquor    LiquorDetail `bson:"liquor"`     // 酒の詳細情報（LiquorDetail）
}

// LiquorDetail 酒の詳細情報
type LiquorDetail struct {
	ID           primitive.ObjectID `bson:"_id"`           // 酒のID
	CategoryID   int                `bson:"category_id"`   // カテゴリID
	CategoryName string             `bson:"category_name"` // カテゴリ名
	Name         string             `bson:"name"`          // 酒の名前
	Description  string             `bson:"description"`   // 説明
	ImageBase64  string             `bson:"image_base64"`  // 画像（base64エンコード）
	ImageURL     string             `bson:"image_url"`     // 画像のURL
	Rate1Users   []string           `bson:"rate1_users"`   // Rate 1 のユーザー
	Rate2Users   []string           `bson:"rate2_users"`   // Rate 2 のユーザー
	Rate3Users   []string           `bson:"rate3_users"`   // Rate 3 のユーザー
	Rate4Users   []string           `bson:"rate4_users"`   // Rate 4 のユーザー
	Rate5Users   []string           `bson:"rate5_users"`   // Rate 5 のユーザー
	UpdatedAt    time.Time          `bson:"updated_at"`    // 更新日時
}

// BoardGroupByRate 各rateごとの掲示板投稿をまとめた構造体
type BoardGroupByRate struct {
	Rate      *int       `json:"_id"`       // 評価（nullも許可）
	Documents []Document `json:"documents"` // 各評価に紐づく投稿
}

// BoardListResponse 返却用の構造体
type BoardListResponse struct {
	GroupedByRate  []BoardGroupByRate `json:"groupedByRate"`   // 評価別の投稿
	RecentComments []Document         `json:"recentDocuments"` // 直近の投稿
}

func (m *BoardModel) ToGraphQL() *graphModel.BoardPost {
	//userはnilの可能性があり、そのままObjectIDを変換して*stringに代入できないので変換
	var userId *string
	if m.UserId != nil {
		id := m.UserId.Hex()
		userId = &id
	}
	return &graphModel.BoardPost{
		ID:         m.ID.Hex(),
		Name:       m.UserName,
		UserID:     userId,
		CategoryID: m.CategoryID,
		LiquorID:   m.LiquorID.Hex(),
		LiquorName: m.LiquorName,
		Text:       m.Text,
		Rate:       m.Rate,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (r *LiquorsRepository) BoardList(ctx context.Context, liquorId primitive.ObjectID) ([]*BoardModel, error) {
	// コレクションからフィルタに一致するドキュメントを取得
	cursor, err := r.boardCollection.Find(ctx, bson.M{LiquorID: liquorId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var liquors []*BoardModel

	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &liquors); err != nil {
		return nil, err
	}

	return liquors, nil
}

// BoardListByUser ユーザーに紐づく掲示板投稿履歴を取得する。評価別および最近のものを取得
func (r *LiquorsRepository) BoardListByUser(ctx context.Context, uId primitive.ObjectID, limit int) (*BoardListResponse, error) {
	log.Println("start")
	pipeline := bson.A{
		bson.M{"$match": bson.M{UserID: uId}}, // フィルタ
		bson.M{"$facet": bson.M{
			"groupedByRate": bson.A{
				bson.M{"$group": bson.M{
					"_id":       "$rate", // rateごとにグループ化
					"documents": bson.M{"$push": "$$ROOT"},
				}},
				bson.M{"$lookup": bson.M{
					"from":         CollectionName,        // 結合するコレクション
					"localField":   "documents.liquor_id", // groupされたドキュメントのliquor_id
					"foreignField": "_id",                 // Liquorコレクションの_idフィールド
					"as":           "liquorDetails",       // 結合結果をliquorDetailsフィールドに格納
				}},
				// liquorDetailsをdocuments内のliquorフィールドとして追加
				bson.M{"$addFields": bson.M{
					"documents": bson.M{
						"$map": bson.M{
							"input": "$documents",
							"as":    "document",
							"in": bson.M{
								"_id":         "$$document._id",
								"category_id": "$$document.category_id",
								"liquor_id":   "$$document.liquor_id",
								"liquor_name": "$$document.liquor_name",
								"text":        "$$document.text",
								"updated_at":  "$$document.updated_at",
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
					"liquorDetails":               0, // liquorDetailsフィールドを除外
					"documents._id":               0,
					"documents.category_id":       0,
					"documents.liquor_id":         0, // documents内のliquor_idを削除
					"documents.liquor_name":       0, // documents内のliquor_nameを削除
					"documents.liquor.version_no": 0, // liquor内のversion_noを削除
				}},
			},
			"recentDocuments": bson.A{
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
					"_id":               0, // liquorDetailsフィールドを除外
					"category_id":       0,
					"liquor_id":         0, // documents内のliquor_idを削除
					"liquor_name":       0, // documents内のliquor_nameを削除
					"user_id":           0,
					"user_name":         0,
					"liquor.version_no": 0, // liquor内のversion_noを削除
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

	//ログ用
	//for cursor.Next(ctx) {
	//	var result bson.M
	//	if err := cursor.Decode(&result); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// JSON形式で整形して出力
	//	jsonData, err := json.MarshalIndent(result, "", "  ")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(string(jsonData))
	//}

	// 返り値の変数を用意
	var result struct {
		GroupedByRate []struct {
			Rate      *int       `bson:"_id"`
			Documents []Document `bson:"documents"`
		} `bson:"groupedByRate"`
		RecentDocuments []Document `bson:"recentDocuments"`
	}

	// クエリ結果を詰め替え
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		// クエリ結果を確認
		fmt.Printf("%+v\n", result)
	}

	log.Println(result)

	// 最終的なレスポンス用構造体に詰め替え
	response := &BoardListResponse{
		RecentComments: result.RecentDocuments,
	}

	// GroupedByRateの詰め替え
	for _, group := range result.GroupedByRate {
		response.GroupedByRate = append(response.GroupedByRate, BoardGroupByRate{
			Rate:      group.Rate,
			Documents: group.Documents,
		})
	}

	return response, nil
}

func (r *LiquorsRepository) BoardGetByUserAndLiquor(ctx context.Context, liquorId primitive.ObjectID, userId primitive.ObjectID, isAllowNotFound bool) (*BoardModel, error) {
	// コレクションからフィルタに一致するドキュメントを取得
	var board *BoardModel
	if err := r.boardCollection.FindOne(ctx, bson.M{LiquorID: liquorId, UserID: userId}).Decode(&board); err != nil {
		if err == mongo.ErrNoDocuments {
			//NotFoundエラーの場合、基本的にはスルーする
			if isAllowNotFound == true {
				return nil, nil
			}
		}
		//それ以外のエラーの場合は普通にエラー
		return nil, err
	}

	return board, nil
}

func (r *LiquorsRepository) BoardInsert(ctx context.Context, board *BoardModel) error {
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
