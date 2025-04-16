package auth

import (
	"backend/db/repository/userRepository"
	"backend/middlewares/customError"
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
func GetId(ctx context.Context) (primitive.ObjectID, *customError.Error) {
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
func GetIdNullable(ctx context.Context) (*primitive.ObjectID, *customError.Error) {
	id, err := GetId(ctx)
	if err != nil {
		return nil, errId()
	}
	if id == primitive.NilObjectID {
		return nil, nil
	}
	return &id, nil
}

func GetIdAndNameNullable(ctx context.Context, ur *userRepository.UsersRepository) (*primitive.ObjectID, *string, *customError.Error) {
	uid, err := GetIdNullable(ctx)
	if err != nil {
		return nil, nil, err
	}
	if uid == nil {
		return nil, nil, nil
	}
	u, err := ur.GetById(ctx, *uid)
	if err != nil {
		return nil, nil, err
	}
	return uid, &u.Name, err
}
