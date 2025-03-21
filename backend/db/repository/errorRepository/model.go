package errorRepository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName = "errors"
	ID             = "id"
	Location       = "location"
	Message        = "message"
	createdAt      = "created_at"
)

// Model 構造体の定義
type Model struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Location  string             `json:"location" bson:"location"`
	Message   string             `json:"message" bson:"message"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
}
