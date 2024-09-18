package liquorRepository

import (
	"backend/graph/graphModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	BoardCollectionName = "liquors_boards"
	LiquorID            = "liquor_id"
	UserId              = "user_id"
	UserName            = "user_name"
)

type BoardModel struct {
	ID        primitive.ObjectID  `bson:"_id"`
	LiquorID  string              `bson:"liquor_id"`
	UserId    *primitive.ObjectID `bson:"user_id"`
	UserName  *string             `bson:"UserName"`
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

func (r *LiquorsRepository) BoardList(ctx context.Context, id string) ([]*BoardModel, error) {
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

func (r *LiquorsRepository) BoardInsert(ctx context.Context, board *BoardModel) error {
	_, err := r.boardCollection.InsertOne(ctx, board)
	if err != nil {
		return err
	}

	return nil
}
