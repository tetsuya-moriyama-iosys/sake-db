package userRepository

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	Register                 = "REPO-USER-001-Register"
	Update                   = "REPO-USER-002-Update"
	GetByEmail               = "REPO-USER-003-GetByEmail"
	GetById                  = "REPO-USER-004-GetById"
	GetByTwitterId           = "REPO-USER-005-GetByTwitterId"
	SetPasswordTokenNotFound = "REPO-USER-006-SetPasswordTokenNotFound"
	SetPasswordToken         = "REPO-USER-007-SetPasswordToken"
	GetByPasswordToken       = "REPO-USER-008-GetByPasswordToken"
	PasswordReset            = "REPO-USER-009-PasswordReset"
)

func errRegister(err error, user *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Register,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      user,
	})
}
func errUpdate(err error, user *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Update,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      user,
	})
}

func errGetByEmail(err error, email string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetByEmail,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      email,
	})
}

func errGetById(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetById,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errGetByTwitterId(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetByTwitterId,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errSetPasswordTokenNotFound(email string, token string) *customError.Error {
	return customError.NewError(errors.New("ユーザーが見つかりません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    SetPasswordTokenNotFound,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      fmt.Printf("email: %s, token: %s", email, token),
	})
}

func errSetPasswordToken(err error, email string, token string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    SetPasswordToken,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("email: %s, token: %s", email, token),
	})
}

func errGetByPasswordToken(err error, token string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    GetByPasswordToken,
		UserMsg:    "有効期限切れです。パスワードリセットURLを再発行してください。",
		Level:      logrus.InfoLevel,
		Input:      token,
	})
}

func errPasswordReset(err error, user Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetByPasswordToken,
		UserMsg:    "パスワードリセットに失敗しました",
		Level:      logrus.InfoLevel,
		Input:      user,
	})
}
