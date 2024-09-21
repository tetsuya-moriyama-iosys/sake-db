package liquorRepository

import (
	"backend/graph/graphModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	BoardCollectionName = "liquors_boards"
	LiquorID            = "liquor_id"
	LiquorName          = "liquor_name"
	UserID              = "user_id"
	UserName            = "user_name"
)

// BoardModel Collectionに挿入するデータ
type BoardModel struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	LiquorID  primitive.ObjectID  `bson:"liquor_id"`
	UserId    *primitive.ObjectID `bson:"user_id"`
	Text      string              `bson:"text"`
	Rate      *int                `bson:"rate"`
	UpdatedAt time.Time           `bson:"updated_at"`
}

// BoardModelWithRelation リレーション込みのモデル(実際に取得してくるデータ)
type BoardModelWithRelation struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty"`
	CategoryID   int                 `bson:"category_id"`
	CategoryName string              `bson:"category_name"`
	LiquorID     primitive.ObjectID  `bson:"liquor_id"`
	LiquorName   string              `bson:"liquor_name"`
	UserId       *primitive.ObjectID `bson:"user_id"`
	UserName     *string             `bson:"user_name"`
	Text         string              `bson:"text"`
	Rate         *int                `bson:"rate"`
	UpdatedAt    time.Time           `bson:"updated_at"`
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
		UserID:    userId,
		LiquorID:  m.LiquorID.Hex(),
		Text:      m.Text,
		Rate:      m.Rate,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *BoardModelWithRelation) ToGraphQL() *graphModel.BoardPost {
	//userはnilの可能性があり、そのままObjectIDを変換して*stringに代入できないので変換
	var userId *string
	if m.UserId != nil {
		id := m.UserId.Hex()
		userId = &id
	}
	return &graphModel.BoardPost{
		ID:           m.ID.Hex(),
		UserName:     m.UserName,
		UserID:       userId,
		CategoryID:   m.CategoryID,
		CategoryName: m.CategoryName,
		LiquorID:     m.LiquorID.Hex(),
		LiquorName:   m.LiquorName,
		Text:         m.Text,
		Rate:         m.Rate,
		UpdatedAt:    m.UpdatedAt,
	}
}

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
