package auth

import (
	"backend/middlewares/customError"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	InvalidToken = "AUTH-001"
	ExpireToken  = "AUTH-002"
	BugToken     = "AUTH-003"
)

func errTokenInvalid(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    InvalidToken,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenExpired(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    ExpireToken,
		UserMsg:    "トークンが期限切れです。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenSomething(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    BugToken,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errMissHeader(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    BugToken,
		UserMsg:    "トークンが見つかりません。",
		Level:      logrus.InfoLevel,
	})
}
func errMissBearer(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    BugToken,
		UserMsg:    "トークンが見つかりません。",
		Level:      logrus.InfoLevel,
	})
}
