package categoriesRepository

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	FindOne            = "REPO-CATEGORY-LOG-001-FindOne"
	FindOneCursor      = "REPO-CATEGORY-LOG-002-FindOneCursor"
	GetLogsByVersionNo = "REPO-CATEGORY-LOG-003-GetLogsByVersionNo"
	StructToBsonM      = "REPO-CATEGORY-LOG-004-StructToBsonM"
	InsertLogOne       = "REPO-CATEGORY-LOG-005-InsertOne"
)

func errFindOne(err error, cId int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    FindOne,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      cId,
	})
}
func errFindOneCursor(err error, cId int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FindOneCursor,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      cId,
	})
}
func errGetLogsByVersionNo(err error, cId int, versionNo int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetLogsByVersionNo,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      fmt.Printf("categoryId:%v,versionNo:%v", cId, versionNo),
	})
}

func errStructToBsonM(err error, cId int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    StructToBsonM,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      cId,
	})
}

func errInsertLogOne(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    InsertLogOne,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}
