package bookmarkService

import (
	"backend/service/userService"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPrimitiveIds(ctx context.Context, targetId string) (primitive.ObjectID, primitive.ObjectID, error) {
	zero := primitive.ObjectID{} // ゼロ値の ObjectID
	targetIdObj, err := primitive.ObjectIDFromHex(targetId)
	if err != nil {
		return zero, zero, err
	}
	uIdObj, err := userService.GetUserId(ctx)
	if err != nil {
		return zero, zero, err
	}
	if targetIdObj == uIdObj {
		return zero, zero, errors.New("自分自身をブックマークできません")
	}
	return uIdObj, targetIdObj, nil
}
