package bookmarkService

import (
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPrimitiveIds(ctx context.Context, targetIdStr string) (primitive.ObjectID, primitive.ObjectID, *customError.Error) {
	zero := primitive.ObjectID{} // ゼロ値の ObjectID
	targetId, err := primitive.ObjectIDFromHex(targetIdStr)
	if err != nil {
		return zero, zero, err
	}
	uId, err := auth.GetId(ctx)
	if err != nil {
		return zero, zero, err
	}
	if targetId == uId {
		return zero, zero, errors.New("自分自身をブックマークできません")
	}
	return uId, targetId, nil
}
