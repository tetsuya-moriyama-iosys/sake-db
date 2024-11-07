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
	CategoryID         = "category_id"
	CategoryName       = "category_name"
	Name               = "name"
	Description        = "description"
	Youtube            = "youtube"
	ImageURL           = "image_url"
	ImageBase64        = "image_base64"
	UpdatedAt          = "updated_at"
	Rate5Users         = "rate5_users"
	Rate4Users         = "rate4_users"
	Rate3Users         = "rate3_users"
	Rate2Users         = "rate2_users"
	Rate1Users         = "rate1_users"
	RandomKey          = "random_key"
	VersionNo          = "version_no"
)

type Model struct {
	ID           primitive.ObjectID `bson:"_id"`
	CategoryID   int                `bson:"category_id"` //カテゴリIDだけは、番号順にソートしたいのでObjectIDではない実装にしている
	CategoryName string             `bson:"category_name"`
	Name         string             `bson:"name"`
	Description  *string            `bson:"description"`
	Youtube      *string            `bson:"youtube"`
	ImageURL     *string            `bson:"image_url"`
	ImageBase64  *string            `bson:"image_base64"`
	Rate5Users   []string           `bson:"rate5_users"`
	Rate4Users   []string           `bson:"rate4_users"`
	Rate3Users   []string           `bson:"rate3_users"`
	Rate2Users   []string           `bson:"rate2_users"`
	Rate1Users   []string           `bson:"rate1_users"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	RandomKey    float64            `bson:"random_key"`
	VersionNo    *int               `bson:"version_no"`
}

func (m *Model) ToGraphQL() *graphModel.Liquor {
	return &graphModel.Liquor{
		ID:           m.ID.Hex(),
		CategoryID:   m.CategoryID,
		CategoryName: m.CategoryName,
		Name:         m.Name,
		Description:  m.Description,
		Youtube:      m.Youtube,
		ImageURL:     m.ImageURL,
		ImageBase64:  m.ImageBase64,
		UpdatedAt:    m.UpdatedAt,
		Rate5Users:   m.Rate5Users,
		Rate4Users:   m.Rate4Users,
		Rate3Users:   m.Rate3Users,
		Rate2Users:   m.Rate2Users,
		Rate1Users:   m.Rate1Users,
		VersionNo:    *m.VersionNo,
	}
}
