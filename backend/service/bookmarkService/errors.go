package bookmarkService

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	GetPrimitiveIdsObjectIDFromHex = "BOOKMARK-SERVICE-001-GetPrimitiveIdsObjectIDFromHex"
	Own                            = "BOOKMARK-SERVICE-002-Own"
	GetBookMarkedListIDFromHex     = "BOOKMARK-SERVICE-003-GetBookMarkedListIDFromHex"
)

func errGetPrimitiveIdsObjectIDFromHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetPrimitiveIdsObjectIDFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errOwn() *customError.Error {
	return customError.NewError(errors.New("自分自身をブックマークできません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    Own,
		UserMsg:    "自分自身をブックマークできません",
		Level:      logrus.InfoLevel,
	})
}

func errGetBookMarkedListIDFromHex(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetBookMarkedListIDFromHex,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}
