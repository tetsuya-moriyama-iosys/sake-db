package myPageService

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	TooShortPassword     = "MYPAGE-SERVICE-001-TooShortPassword"
	GenerateFromPassword = "MYPAGE-SERVICE-002-GenerateFromPassword"
)

func errTooShortPassword() *customError.Error {
	return customError.NewError(errors.New("パスワードが短いです"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    TooShortPassword,
		UserMsg:    "パスワードが短いです",
		Level:      logrus.InfoLevel,
	})
}
func errGenerateFromPassword(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GenerateFromPassword,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}
