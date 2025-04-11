package authService

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"fmt"
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

const (
	NotFoundPassOrMail = "LOGIN-001"
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

func errLogin() *customError.Error {
	return customError.NewError(errors.New("メールアドレスもしくはパスワードが間違っています。"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundPassOrMail,
		UserMsg:    "メールアドレスもしくはパスワードが間違っています。",
		Level:      logrus.InfoLevel,
	})
}

const (
	GenerateFromPassword = "AUTH-PASSWORD-RESET-001-GenerateFromPassword"
	SendPasswordReset    = "AUTH-PASSWORD-RESET-002-SendPasswordReset"
)

func errGenerateFromPassword(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GenerateFromPassword,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errSendPasswordReset(err error, email string, token string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    SendPasswordReset,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("email: %s, token: %s", email, token),
	})
}
