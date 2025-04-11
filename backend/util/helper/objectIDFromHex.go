package helper

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// ObjectIDFromHex 雑なカスタムエラーを返しても良い時のヘルパ
func ObjectIDFromHex(id string) (primitive.ObjectID, *customError.Error) {
	// 文字列をObjectIDに変換
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, customError.NewError(err, customError.Params{
			StatusCode: http.StatusBadRequest,
			ErrCode:    "HELPER-ObjectIDFromHex",
			UserMsg:    errorMsg.SERVER,
			Level:      logrus.InfoLevel,
			Input:      id,
		})
	}
	return objectID, nil
}
