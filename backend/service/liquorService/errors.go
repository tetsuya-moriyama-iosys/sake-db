package liquorService

import (
	"backend/db/repository/liquorRepository"
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	GetLiquorIdHex              = "LIQUOR-SERVICE-001-GetLiquorIdHex"
	GetLiquorId                 = "LIQUOR-SERVICE-002-GetLiquorId"
	PostBoardObjectIDFromHex    = "LIQUOR-SERVICE-003-PostBoardIdHex"
	PostBoardErr                = "LIQUOR-SERVICE-004-PostBoard"
	GetLiquorHistoriesIDFromHex = "LIQUOR-SERVICE-005-GetLiquorHistoriesIdHex"
	GetBoardFromHex             = "LIQUOR-SERVICE-006-GetBoardFromHex"
	GetMyBoardFromHex           = "LIQUOR-SERVICE-007-GetMyBoardFromHex "
)

func errGetLiquorIdHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetLiquorIdHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errGetLiquorId(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetLiquorId,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errPostBoardObjectIDFromHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    PostBoardObjectIDFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errPostBoard(err error, model *liquorRepository.BoardModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    PostBoardErr,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      model,
	})
}

func errGetLiquorHistoriesIDFromHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorHistoriesIDFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errGetBoardFromHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetBoardFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errGetMyBoard(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetMyBoardFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}
