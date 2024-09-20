package liquorRepository

import (
	"backend/graph/graphModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	BoardCollectionName = "liquors_boards"
	LiquorID            = "liquor_id"
	UserID              = "user_id"
	UserName            = "user_name"
)

type BoardModel struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	LiquorID  primitive.ObjectID  `bson:"liquor_id"`
	UserId    *primitive.ObjectID `bson:"user_id"`
	UserName  *string             `bson:"user_name"`
	Text      string              `bson:"text"`
	Rate      *int                `bson:"rate"`
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
}

// BoardGroupByRate 各rateごとの掲示板投稿をまとめた構造体
type BoardGroupByRate struct {
	Rate      *int         `json:"rate"`      // 評価（nullも許可）
	Documents []BoardModel `json:"documents"` // 各評価に紐づく投稿
}

// BoardListResponse 返却用の構造体
type BoardListResponse struct {
	GroupedByRate   []BoardGroupByRate `json:"grouped_by_rate"`  // 評価別の投稿
	RecentDocuments []BoardModel       `json:"recent_documents"` // 直近の投稿
}

func (m *BoardModel) ToGraphQL() *graphModel.BoardPost {
	//userはnilの可能性があり、そのままObjectIDを変換して*stringに代入できないので変換
	var userId *string
	if m.UserId != nil {
		id := m.UserId.Hex()
		userId = &id
	}
	return &graphModel.BoardPost{
		ID:        m.ID.Hex(),
		Name:      m.UserName,
		UserID:    userId,
		LiquorID:  m.LiquorID.Hex(),
		Text:      m.Text,
		Rate:      m.Rate,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
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
func (r *LiquorsRepository) BoardListByUser(ctx context.Context, UserId primitive.ObjectID, limit int) (*BoardListResponse, error) {
	pipeline := bson.A{
		// 1. user_idがidであるドキュメントをフィルタリング
		bson.M{"$match": bson.M{"user_id": UserId}},

		bson.M{"$facet": bson.M{
			// rate別にグルーピング
			"groupedByRate": bson.A{
				bson.M{"$match": bson.M{"rate": bson.M{"$in": []interface{}{1, 2, 3, 4, 5, nil}}}},
				bson.M{"$group": bson.M{
					"_id":       "$rate",
					"documents": bson.M{"$push": "$$ROOT"},
				}},
			},
			// 直近10件の取得
			"recentDocuments": bson.A{
				bson.M{"$sort": bson.M{"createdAt": -1}}, // createdAtの降順にソート
				bson.M{"$limit": limit},                  // 直近n件を取得
			},
		}},
	}

	// MongoDBの集計クエリを実行
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 返り値の変数を用意
	var result struct {
		GroupedByRate []struct {
			Rate      *int         `bson:"_id"`
			Documents []BoardModel `bson:"documents"`
		} `bson:"groupedByRate"`
		RecentDocuments []BoardModel `bson:"recentDocuments"`
	}

	// クエリ結果を詰め替え
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	}

	// 最終的なレスポンス用構造体に詰め替え
	response := &BoardListResponse{
		RecentDocuments: result.RecentDocuments,
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
