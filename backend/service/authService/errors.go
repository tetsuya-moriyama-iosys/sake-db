package authService

import (
	"backend/middlewares/customError"
	"errors"
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
	})
}

func errTokenExpired() *customError.Error {
	return customError.NewError(errors.New("expired refresh token"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenExpired,
		UserMsg:    "トークンが期限切れです。",
	})
}

func errTokenInvalid() *customError.Error {
	return customError.NewError(errors.New("invalid refresh token"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenInvalid,
		UserMsg:    "トークンが不正です。",
	})
}

func errInvalidClimes() *customError.Error {
	return customError.NewError(errors.New("invalid claims"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenInvalidClimes,
		UserMsg:    "トークンが不正です。",
	})
}

func errRefreshToken(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    RefreshTokenInvalid,
		UserMsg:    "リフレッシュトークンが期限切れです。",
	})
}
