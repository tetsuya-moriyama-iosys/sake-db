package bookmarkService

import (
	"backend/db/repository/bookmarkRepository"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPrimitiveIds(ctx context.Context, targetIdStr string) (primitive.ObjectID, primitive.ObjectID, *customError.Error) {
	zero := primitive.ObjectID{} // ゼロ値の ObjectID
	targetId, err := primitive.ObjectIDFromHex(targetIdStr)
	if err != nil {
		return zero, zero, errGetPrimitiveIdsObjectIDFromHex(err, targetIdStr)
	}
	uId, cErr := auth.GetId(ctx)
	if cErr != nil {
		return zero, zero, cErr
	}
	if targetId == uId {
		return zero, zero, errOwn()
	}
	return uId, targetId, nil
}

func GetBookMarkedList(ctx context.Context, r bookmarkRepository.BookMarkRepository, id string) ([]*bookmarkRepository.BookMarkListUser, *customError.Error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errGetBookMarkedListIDFromHex(err, id)
	}
	return r.BookmarkedList(ctx, idObj)
}
