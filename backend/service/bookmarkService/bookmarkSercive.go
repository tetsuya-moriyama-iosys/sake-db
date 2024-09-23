package bookmarkService

import (
	"backend/service/userService"
	"context"
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
	return uIdObj, targetIdObj, nil
}
