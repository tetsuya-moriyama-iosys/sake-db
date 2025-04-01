package liquorRepository

import (
	"backend/middlewares/customError"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	GetLogsById = "REPO-LIQUOR_LOG-001-GetLogsById"
	Cursor      = "REPO-LIQUOR_LOG-002-Cursor"

	GetLogsByVer = "REPO-LIQUOR_LOG-003-GetLogsByVer"

	ToBsonForInsert = "REPO-LIQUOR_LOG-004-ToBsonForInsert"
	InsertLogOne    = "REPO-LIQUOR_LOG-005-InsertOne"
)

func errGetLogsById(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLogsById,
		UserMsg:    "バージョンログ取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errCursor(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Cursor,
		UserMsg:    "バージョンログ取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errGetLogsByVer(err error, id string, versionNo int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLogsByVer,
		UserMsg:    "バージョン取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("id:%v,version:%v", id, versionNo),
	})
}
func errToBsonForInsert(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ToBsonForInsert,
		UserMsg:    "不明なエラーが発生しました",
		Level:      logrus.ErrorLevel,
	})
}

func errInsertLogOne(err error, oldLiquor *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    InsertLogOne,
		UserMsg:    "データの追加に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      oldLiquor,
	})
}
