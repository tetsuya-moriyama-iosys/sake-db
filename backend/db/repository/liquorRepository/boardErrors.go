package liquorRepository

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	GetList                 = "REPO-LIQUOR-BOARD-001-GetList"
	GetListDecode           = "REPO-LIQUOR-BOARD-002-GetListDecode"
	BoardListByUser         = "REPO-LIQUOR-BOARD-003-BoardListByUser"
	BoardListByUserDecode   = "REPO-LIQUOR-BOARD-004-BoardListByUserDecode"
	BoardGetByUserAndLiquor = "REPO-LIQUOR-BOARD-005-BoardGetByUserAndLiquor"
	BoardInsertGuest        = "REPO-LIQUOR-BOARD-006-BoardInsertGuest"
	BoardUpsert             = "REPO-LIQUOR-BOARD-007-BoardUpsert"
)

func errGetList(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetList,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errGetListDecode(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetListDecode,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errBoardListByUser(err error, uId primitive.ObjectID, limit int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BoardListByUser,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("uId: %v, limit: %d", uId, limit),
	})
}
func errBoardListByUserDecode(err error, uId primitive.ObjectID, limit int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BoardListByUserDecode,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("uId: %v, limit: %d", uId, limit),
	})
}

func errBoardGetByUserAndLiquor(err error, liquorId primitive.ObjectID, userId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BoardGetByUserAndLiquor,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("liquorId: %v, userId: %v", liquorId, userId),
	})
}

func errBoardInsertGuest(err error, board *BoardModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BoardInsertGuest,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      board,
	})
}

func errBoardUpsert(err error, board *BoardModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BoardUpsert,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      board,
	})
}
