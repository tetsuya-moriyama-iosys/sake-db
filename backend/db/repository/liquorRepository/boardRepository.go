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
		Text:      m.Text,
		Rate:      m.Rate,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r *LiquorsRepository) BoardList(ctx context.Context, id primitive.ObjectID) ([]*BoardModel, error) {
	// コレクションからフィルタに一致するドキュメントを取得
	cursor, err := r.boardCollection.Find(ctx, bson.M{"liquor_id": id})
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

func (r *LiquorsRepository) BoardGetByUserAndLiquor(ctx context.Context, liquorId primitive.ObjectID, userId primitive.ObjectID) (*BoardModel, error) {
	// コレクションからフィルタに一致するドキュメントを取得
	var board *BoardModel
	if err := r.boardCollection.FindOne(ctx, bson.M{LiquorID: liquorId, UserID: userId}).Decode(&board); err != nil {
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
