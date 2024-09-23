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

// BookMarkListUser ユーザーページのブックマークリストの構造体
type BookMarkListUser struct {
	UserId   primitive.ObjectID `json:"userId" bson:"user_id"`
	UserName string             `json:"userName" bson:"user_name"`
}

type BookMarkList []*BookMarkListUser

func (l BookMarkList) ToGraphQL() []*graphModel.BookMarkListUser {
	var result []*graphModel.BookMarkListUser
	for _, b := range l {
		result = append(result, &graphModel.BookMarkListUser{
			UserID:    b.UserId.Hex(),
			Name:      b.UserName,
			CreatedAt: b.UserId.Timestamp(),
		})
	}
	return result
}
