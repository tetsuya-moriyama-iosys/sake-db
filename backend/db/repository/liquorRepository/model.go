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
	Rate5Users         = "rate5_users"
	Rate4Users         = "rate4_users"
	Rate3Users         = "rate3_users"
	Rate2Users         = "rate2_users"
	Rate1Users         = "rate1_users"
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
	Rate5Users   []string           `bson:"rate5_users"`
	Rate4Users   []string           `bson:"rate4_users"`
	Rate3Users   []string           `bson:"rate3_users"`
	Rate2Users   []string           `bson:"rate2_users"`
	Rate1Users   []string           `bson:"rate1_users"`
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
		Rate5Users:   m.Rate5Users,
		Rate4Users:   m.Rate4Users,
		Rate3Users:   m.Rate3Users,
		Rate2Users:   m.Rate2Users,
		Rate1Users:   m.Rate1Users,
		VersionNo:    *m.VersionNo,
	}
}

func (m *Model) ToGraphQLSimple() *graphModel.LiquorSimple {
	return &graphModel.LiquorSimple{
		ID:           m.ID.Hex(),
		CategoryID:   m.CategoryID,
		CategoryName: m.CategoryName,
		Name:         m.Name,
		ImageBase64:  m.ImageBase64,
	}
}
