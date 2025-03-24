package authService

import (
	"backend/middlewares/customError"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	TokenNotFound       = "TOKEN-001"
	TokenExpired        = "TOKEN-002"
	TokenInvalid        = "TOKEN-003"
	TokenInvalidClimes  = "TOKEN-004"
	RefreshTokenInvalid = "TOKEN-005"
)

func errTokenNotFound() *customError.Error {
	return customError.NewError(errors.New("refresh token not found"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenNotFound,
		UserMsg:    "トークンがありません。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenExpired() *customError.Error {
	return customError.NewError(errors.New("expired refresh token"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenExpired,
		UserMsg:    "トークンが期限切れです。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenInvalid() *customError.Error {
	return customError.NewError(errors.New("invalid refresh token"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenInvalid,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errInvalidClimes() *customError.Error {
	return customError.NewError(errors.New("invalid claims"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenInvalidClimes,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errRefreshToken(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    RefreshTokenInvalid,
		UserMsg:    "リフレッシュトークンが期限切れです。",
		Level:      logrus.InfoLevel,
	})
}
