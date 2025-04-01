package categoryPost

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	ParseFailInput = "CATEGORY-POST-001-ParseFailInput"
	InvalidInput   = "CATEGORY-POST-002-InvalidInput"
	InvalidParent  = "CATEGORY-POST-003-InvalidParent"
	InvalidVersion = "CATEGORY-POST-004-InvalidVersion"
	InvalidFile    = "CATEGORY-POST-005-InvalidFile"
)

func errInvalidInput(c *gin.Context, err error) *customError.Error {
	raw, getRawErr := c.GetRawData()
	if getRawErr != nil {
		return customError.NewError(err, customError.Params{
			StatusCode: http.StatusBadRequest,
			ErrCode:    ParseFailInput,
			UserMsg:    "入力値が不正です",
			Level:      logrus.ErrorLevel,
		})
	}
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidInput,
		UserMsg:    "入力値が不正です",
		Level:      logrus.ErrorLevel,
		Input:      raw,
	})
}

func errInvalidParent(input RequestData) *customError.Error {
	return customError.NewError(errors.New("自身または子カテゴリを親とすることはできません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidParent,
		UserMsg:    "自身または子カテゴリを親とすることはできません",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errInvalidVersion(input RequestData) *customError.Error {
	return customError.NewError(errors.New("自身または子カテゴリを親とすることはできません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidVersion,
		UserMsg:    errorMsg.VERSION,
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errInvalidFile(err error, input RequestData) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidFile,
		UserMsg:    "ファイルが不正です",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}
