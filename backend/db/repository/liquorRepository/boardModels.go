package liquorRepository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	BoardCollectionName = "liquors_boards"
	Rate                = "rate"
	LiquorID            = "liquor_id"
	LiquorName          = "liquor_name"
	UserID              = "user_id"
	UserName            = "user_name"
)

// BoardModel Collectionに挿入するデータ
type BoardModel struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	LiquorID  primitive.ObjectID  `bson:"liquor_id"`
	UserId    *primitive.ObjectID `bson:"user_id"`
	Text      string              `bson:"text"`
	Rate      *int                `bson:"rate"`
	UpdatedAt time.Time           `bson:"updated_at"`
}

// BoardModelWithRelation リレーション込みのモデル(実際に取得してくるデータ)
type BoardModelWithRelation struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty"`
	CategoryID   int                 `bson:"category_id"`
	CategoryName string              `bson:"category_name"`
	LiquorID     primitive.ObjectID  `bson:"liquor_id"`
	LiquorName   string              `bson:"liquor_name"`
	UserId       *primitive.ObjectID `bson:"user_id"`
	UserName     *string             `bson:"user_name"`
	Text         string              `bson:"text"`
	Rate         *int                `bson:"rate"`
	UpdatedAt    time.Time           `bson:"updated_at"`
}

// Post 各投稿の詳細
type Post struct {
	ID        primitive.ObjectID `bson:"_id"`        // 投稿内容
	Text      string             `bson:"text"`       // 投稿内容
	Rate      *int               `bson:"rate"`       //RecentPostsと共用するために追加(GroupedByRateのPosts内では、更にリレーションを取る必要はないので常にnilになる)
	UpdatedAt time.Time          `bson:"updated_at"` // 更新日時
	Liquor    LiquorDetail       `bson:"liquor"`     // 酒の詳細情報（LiquorDetail）
}

// LiquorDetail 酒の詳細情報
type LiquorDetail struct {
	ID           primitive.ObjectID `bson:"_id"`           // 酒のID
	CategoryID   int                `bson:"category_id"`   // カテゴリID
	CategoryName string             `bson:"category_name"` // カテゴリ名
	Name         string             `bson:"name"`          // 酒の名前
	Description  string             `bson:"description"`   // 説明
	ImageBase64  *string            `bson:"image_base64"`  // 画像（base64エンコード）
	ImageURL     string             `bson:"image_url"`     // 画像のURL
	Rate1Users   []string           `bson:"rate1_users"`   // Rate 1 のユーザー
	Rate2Users   []string           `bson:"rate2_users"`   // Rate 2 のユーザー
	Rate3Users   []string           `bson:"rate3_users"`   // Rate 3 のユーザー
	Rate4Users   []string           `bson:"rate4_users"`   // Rate 4 のユーザー
	Rate5Users   []string           `bson:"rate5_users"`   // Rate 5 のユーザー
	UpdatedAt    time.Time          `bson:"updated_at"`    // 更新日時
}

// BoardGroupByRate 各rateごとの掲示板投稿をまとめた構造体
type BoardGroupByRate struct {
	Rate  *int   `bson:"_id"`   // 評価（nullも許可）
	Posts []Post `bson:"posts"` // 各評価に紐づく投稿
}

// BoardListResponse 返却用の構造体
type BoardListResponse struct {
	GroupedByRate []BoardGroupByRate `bson:"groupedByRate"` // 評価別の投稿
	RecentPosts   []Post             `bson:"recentPosts"`   // 直近の投稿
}
