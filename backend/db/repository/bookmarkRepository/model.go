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

// BookMarkListUser ユーザーページのブックマークリストの構造体
type BookMarkListUser struct {
	UserId    primitive.ObjectID `json:"userId" bson:"user_id"`
	UserName  string             `json:"userName" bson:"user_name"`
	CreatedAt time.Time          `bson:"created_at"`
}

type BookMarkList []*BookMarkListUser

type Recommend struct {
	Comment string          `json:"comment"`
	liquor  RecommendLiquor `json:"liquor"`
}

type RecommendLiquor struct {
	ID           primitive.ObjectID `json:"_id"`
	CategoryID   string             `json:"categoryId"`
	CategoryName string             `json:"categoryName"`
	ImageBase64  string             `json:"imageBase64"`
	ImageBase64  string             `json:"imageBase64"`
}

type RecommendUser struct{}

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
