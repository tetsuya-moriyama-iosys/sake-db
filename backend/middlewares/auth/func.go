package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ContextKey ユーザーIDを保存するためのコンテキストキー
const userContextKey = "user"

func setId(ctx context.Context, id primitive.ObjectID) context.Context {
	return context.WithValue(ctx, userContextKey, id)
}

// GetId コンテキストからユーザーIDを取得する(型安全にするため定義)
func GetId(ctx context.Context) (primitive.ObjectID, error) {
	rawId := ctx.Value(userContextKey)
	if rawId == nil {
		return primitive.NilObjectID, errId()
	}
	id := rawId.(primitive.ObjectID)
	if id == primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("unauthorized")
	}
	return id, nil
}

// GetIdNullable nil許容の場合
func GetIdNullable(ctx context.Context) *primitive.ObjectID {
	id, _ := GetId(ctx)
	if id == primitive.NilObjectID {
		return nil
	}
	return &id
}
