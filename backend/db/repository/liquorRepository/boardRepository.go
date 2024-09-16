package liquorRepository

import (
	"context"
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
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
}

func (r *LiquorsRepository) BoardInsert(ctx context.Context, board *BoardModel) error {
	_, err := r.boardCollection.InsertOne(ctx, board)
	if err != nil {
		return err
	}

	return nil
}
