package liquorModel

import (
	"backend/graph/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName = "liquors"
	ID             = "_id"
	CategoryID     = "category_id"
	CategoryName   = "category_name"
	Name           = "name"
	Description    = "description"
	ImageURL       = "image_url"
	ImageBase64    = "image_base64"
	CreatedAt      = "created_at"
	UpdatedAt      = "updated_at"
)

type Schema struct {
	ID           primitive.ObjectID `bson:"_id"`
	CategoryID   int                `bson:"category_id"`
	CategoryName string             `bson:"category_name"`
	Name         string             `bson:"name"`
	Description  *string            `bson:"description"`
	ImageURL     *string            `bson:"image_url"`
	ImageBase64  *string            `bson:"image_base64"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

func (ldb *Schema) ToGraphQL() *model.Liquor {
	return &model.Liquor{
		ID:           ldb.ID.Hex(),
		CategoryID:   ldb.CategoryID,
		CategoryName: ldb.CategoryName,
		Name:         ldb.Name,
		Description:  ldb.Description,
		ImageURL:     ldb.ImageURL,
		ImageBase64:  ldb.ImageBase64,
		CreatedAt:    ldb.CreatedAt,
		UpdatedAt:    ldb.UpdatedAt,
	}
}

func FromGraphQL(l *model.Liquor) *Schema {
	objectID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		fmt.Println("Invalid ObjectID string")
	}

	return &Schema{
		ID:          objectID,
		CategoryID:  l.CategoryID,
		Name:        l.Name,
		Description: l.Description,
		ImageURL:    l.ImageURL,
		ImageBase64: l.ImageBase64,
		CreatedAt:   l.CreatedAt,
		UpdatedAt:   l.UpdatedAt,
	}
}
