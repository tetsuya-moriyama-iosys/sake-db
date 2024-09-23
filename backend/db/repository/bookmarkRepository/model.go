package bookmarkRepository

import (
	"backend/graph/graphModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionName     = "bookmarks"
	ID                 = "_id"
	USER_ID            = "user_id"
	USER_NAME          = "user_name"
	BOOKMARKED_USER_ID = "bookmarked_user_id"
	UPDATED_AT         = "updated_at"
)

type Model struct {
	UserId           primitive.ObjectID `bson:"user_id"`
	BookmarkedUserId primitive.ObjectID `bson:"bookmarked_user_id"`
}

// BookMarkListUser 構造体の定義
type BookMarkListUser struct {
	//ID          int       `json:"id" bson:"_id"`
	UserId   string `json:"userId" bson:"user_id"`
	UserName string `json:"userName" bson:"user_name"`
	//BookmarkedUserId      string      `json:"bookmarkedUserId" bson:"bookmarked_user_id"`
}

func (m *BookMarkListUser) ToGraphQL() *graphModel.BookMarkListUser {
	return &graphModel.BookMarkListUser{
		UserID: m.UserId,
		Name:   m.UserName,
	}
}
