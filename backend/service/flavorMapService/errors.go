package flavorMapService

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	NotFoundMstData        = "FLAVOR-SERVICE-001-GetFlavorMasterData"
	NotFound               = "FLAVOR-SERVICE-002-NotFound"
	Cursor                 = "FLAVOR-SERVICE-003-Cursor"
	InsertOne              = "FLAVOR-SERVICE-004-InsertOne"
	PostFlavorMapIdFromHex = "FLAVOR-SERVICE-005-PostFlavorMapIdFromHex"
)

func errNotFoundMstData(id primitive.ObjectID) *customError.Error {
	return customError.NewError(errors.New("フレーバーマップが存在しません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    NotFoundMstData,
		UserMsg:    "フレーバーマップが存在しません",
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errNotFound(err error, lId primitive.ObjectID, cId int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    NotFound,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      fmt.Printf("lId: %s, cId: %d", lId.Hex(), cId),
	})
}
func errCursor(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Cursor,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errInsertOne(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    InsertOne,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errPostFlavorMapIdFromHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    PostFlavorMapIdFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}
