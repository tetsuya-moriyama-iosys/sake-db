package liquorRepository

import (
	"backend/graph/graphModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName     = "liquors"
	LogsCollectionName = "liquors_logs"
	ID                 = "_id"
	LogID              = "id"
	CategoryID         = "category_id"
	CategoryName       = "category_name"
	Name               = "name"
	Description        = "description"
	ImageURL           = "image_url"
	ImageBase64        = "image_base64"
	CreatedAt          = "created_at"
	UpdatedAt          = "updated_at"
	VersionNo          = "version_no"
)

type Model struct {
	ID           primitive.ObjectID `bson:"_id"`
	CategoryID   int                `bson:"category_id"`
	CategoryName string             `bson:"category_name"`
	Name         string             `bson:"name"`
	Description  *string            `bson:"description"`
	ImageURL     *string            `bson:"image_url"`
	ImageBase64  *string            `bson:"image_base64"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	VersionNo    *int               `bson:"version_no"`
}

func (m *Model) ToGraphQL() *graphModel.Liquor {
	return &graphModel.Liquor{
		ID:           m.ID.Hex(),
		CategoryID:   m.CategoryID,
		CategoryName: m.CategoryName,
		Name:         m.Name,
		Description:  m.Description,
		ImageURL:     m.ImageURL,
		ImageBase64:  m.ImageBase64,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		VersionNo:    *m.VersionNo,
	}
}

//func FromGraphQL(l *graphModel.Liquor) *Model {
//	objectID, err := primitive.ObjectIDFromHex(l.ID)
//	if err != nil {
//		fmt.Println("Invalid ObjectID string")
//	}
//
//	return &Model{
//		ID:          objectID,
//		CategoryID:  l.CategoryID,
//		Name:        l.Name,
//		Description: l.Description,
//		ImageURL:    l.ImageURL,
//		ImageBase64: l.ImageBase64,
//		CreatedAt:   l.CreatedAt,
//		UpdatedAt:   l.UpdatedAt,
//		VersionNo:   l.VersionNo,
//	}
//}
