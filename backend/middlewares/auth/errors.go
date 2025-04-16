package auth

import (
	"backend/middlewares/customError"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	InvalidToken   = "AUTH-001"
	ExpireToken    = "AUTH-002"
	BugToken       = "AUTH-003"
	NotFoundToken  = "AUTH-004"
	NotFoundBearer = "AUTH-005"
	ErrorId        = "AUTH-006"
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

func errMissHeader() *customError.Error {
	return customError.NewError(errors.New("authorization header is missing"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundToken,
		UserMsg:    "トークンが見つかりません。",
		Level:      logrus.InfoLevel,
	})
}

func errMissBearer() *customError.Error {
	return customError.NewError(errors.New("authorization token is missing"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundBearer,
		UserMsg:    "トークンが見つかりません。",
		Level:      logrus.InfoLevel,
	})
}

func errId() *customError.Error {
	return customError.NewError(errors.New("unauthorized"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    ErrorId,
		UserMsg:    "ユーザーIDがが見つかりません。",
		Level:      logrus.InfoLevel,
	})
}
