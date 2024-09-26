package bookmarkRepository

import (
	"backend/graph/graphModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName   = "bookmarks"
	UserID           = "user_id"
	UserName         = "user_name"
	BookmarkedUserId = "bookmarked_user_id"
	CreatedAT        = "created_at"
)

type Model struct {
	UserId           primitive.ObjectID `bson:"user_id"`
	BookmarkedUserId primitive.ObjectID `bson:"bookmarked_user_id"`
}

// BookMarkList ブックマークリスト
type BookMarkList []*BookMarkListUser

// BookMarkListUser ユーザーページのブックマークリストの構造体
type BookMarkListUser struct {
	UserId    primitive.ObjectID `json:"userId" bson:"user_id"`
	UserName  string             `json:"userName" bson:"user_name"`
	CreatedAt time.Time          `bson:"created_at"`
}

// RecommendList リコメンドリスト
type RecommendList []*Recommend

type Recommend struct {
	Rate      int             `bson:"rate"`
	Comment   string          `bson:"comment"`
	Liquor    RecommendLiquor `bson:"liquor"`
	User      RecommendUser   `bson:"user_info"`
	UpdatedAt time.Time       `bson:"updated_at"`
}

type RecommendLiquor struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	CategoryID   int                `bson:"category_id"`
	CategoryName string             `bson:"category_name"`
	ImageBase64  *string            `bson:"image_base64"`
	Description  string             `bson:"description"`
}

type RecommendUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	ImageBase64 *string            `bson:"image_base64"`
}

func (l BookMarkList) ToGraphQL() []*graphModel.BookMarkListUser {
	var result []*graphModel.BookMarkListUser
	for _, b := range l {
		result = append(result, &graphModel.BookMarkListUser{
			UserID:    b.UserId.Hex(),
			Name:      b.UserName,
			CreatedAt: b.CreatedAt,
		})
	}
	return result
}

func (l RecommendList) ToGraphQL() []*graphModel.Recommend {
	var result []*graphModel.Recommend
	for _, m := range l {
		result = append(result, &graphModel.Recommend{
			Rate:      m.Rate,
			Comment:   m.Comment,
			UpdatedAt: m.UpdatedAt,
			Liquor:    m.Liquor.ToGraphQL(),
			User:      m.User.ToGraphQL(),
		})
	}
	return result
}

func (l RecommendLiquor) ToGraphQL() *graphModel.RecommendLiquor {
	return &graphModel.RecommendLiquor{
		ID:           l.ID.Hex(),
		Name:         l.Name,
		CategoryID:   l.CategoryID,
		CategoryName: l.CategoryName,
		ImageBase64:  l.ImageBase64,
		Description:  l.Description,
	}
}
func (u RecommendUser) ToGraphQL() *graphModel.RecommendUser {
	return &graphModel.RecommendUser{
		ID:          u.ID.Hex(),
		Name:        u.Name,
		ImageBase64: u.ImageBase64,
	}
}
