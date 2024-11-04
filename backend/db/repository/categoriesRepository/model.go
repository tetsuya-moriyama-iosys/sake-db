package categoriesRepository

import (
	"backend/graph/graphModel"
	"time"
)

const (
	CollectionName     = "categories"
	LogsCollectionName = "categories_logs"
	ID                 = "id"
	Name               = "name"
	ImageURL           = "image_url"
	ImageBase64        = "image_base64"
	Description        = "description"
	Parent             = "parent"
	VersionNo          = "version_no"
	Order              = "order"
	UpdatedAt          = "updated_at"
	Authorized         = "authorized"
)

// Model 構造体の定義
type Model struct {
	ID          int     `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Parent      *int    `json:"parent" bson:"parent"`
	Description *string `bson:"description"`
	ImageURL    *string `bson:"image_url"`
	ImageBase64 *string `bson:"image_base64"`
	//Authorized  *bool      `bson:"authorized"`                  //カテゴリ移動不可フラグ
	VersionNo *int      `json:"versionNo" bson:"version_no"` //手動で追加したカテゴリはversionNoが存在しない可能性がある
	Children  []*Model  `json:"children,omitempty"`          // 子カテゴリはDBに保存されないため、bsonタグは不要
	Order     *int      `bson:"order"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (s *Model) ToGraphQL() *graphModel.Category {
	// 子カテゴリを再帰的に変換
	var children []*graphModel.Category
	if len(s.Children) > 0 {
		children = make([]*graphModel.Category, len(s.Children))
		for i, child := range s.Children {
			children[i] = child.ToGraphQL() // 再帰的にToGraphQLを呼び出す
		}
	}

	return &graphModel.Category{
		ID:          s.ID,
		Name:        s.Name,
		Parent:      s.Parent,
		Description: s.Description,
		ImageURL:    s.ImageURL,
		ImageBase64: s.ImageBase64,
		VersionNo:   s.VersionNo,
		UpdatedAt:   &s.UpdatedAt,
		//Authorized:  s.Authorized,
		Children: children, // 変換後の子カテゴリを設定
	}
}
