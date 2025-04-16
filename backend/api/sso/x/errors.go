package x

import (
	"backend/middlewares/customError"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	GenerateRandomKey     = "X-001-GenerateRandomKey"
	ParseURL              = "X-002-ParseURL"
	MissCode              = "X-003-MissCode"
	MissExchangeToken     = "X-004-MissExchangeToken"
	GetUserInfoFail       = "X-005-GetUserInfoFail"
	GetUserInfoBadStatus  = "X-006-GetUserInfoBadStatus"
	ReadResponseBody      = "X-007-ReadResponseBody"
	UnMarshalResponseBody = "X-008-UnMarshalResponseBody"
	GetData               = "X-009-GetData"
)

func errInvalidInput(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GenerateRandomKey,
		UserMsg:    "キー生成に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}
func errParseURL(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ParseURL,
		UserMsg:    "キー生成に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}
func errMissCode() *customError.Error {
	return customError.NewError(errors.New("missing code"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    MissCode,
		UserMsg:    "サーバーでエラーが発生しました",
		Level:      logrus.FatalLevel,
	})
}
func errMissExchangeToken(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    MissExchangeToken,
		UserMsg:    "サーバーでエラーが発生しました",
		Level:      logrus.ErrorLevel,
	})
}
func errGetUserInfo(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetUserInfoFail,
		UserMsg:    "ユーザー情報取得に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}

func errGetUserInfoBadStatus(errBody []byte) *customError.Error {
	return customError.NewError(errors.New("failed to fetch user info"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetUserInfoBadStatus,
		UserMsg:    "ユーザー情報取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      errBody,
	})
}
func errReadResponseBody(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ReadResponseBody,
		UserMsg:    "サーバーのレスポンスが不正です",
		Level:      logrus.ErrorLevel,
	})
}

func errUnMarshalResponseBody(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    UnMarshalResponseBody,
		UserMsg:    "サーバーのレスポンスが不正です",
		Level:      logrus.ErrorLevel,
	})
}
func errGetData(input map[string]interface{}) *customError.Error {
	return customError.NewError(errors.New("failed reading data"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetData,
		UserMsg:    "サーバーのレスポンスが不正です",
		Level:      logrus.ErrorLevel,
		Input:      input,
	})
}
