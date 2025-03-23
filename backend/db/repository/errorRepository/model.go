package errorRepository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName = "errors"
	ID             = "id"
	UserId         = "user_id"
	Code           = "code"
	Location       = "location"
	Message        = "message"
	Input          = "input"
	createdAt      = "created_at"
)

// Model 構造体の定義
type Model struct {
	ID        primitive.ObjectID  `json:"id" bson:"_id"`
	Code      string              `json:"code" bson:"code"`
	UserId    *primitive.ObjectID `json:"userId" bson:"user_id"`
	Location  string              `json:"location" bson:"location"`
	Message   string              `json:"message" bson:"message"`
	Input     string              `json:"input" bson:"input"`
	CreatedAt time.Time           `json:"createdAt" bson:"created_at"`
}
