package liquorRepository

import (
	"backend/middlewares/customError"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	GetLiquorById                 = "REPO-LIQUOR-001-GetLiquorById"
	GetLiquorByName               = "REPO-LIQUOR-002-GetLiquorByName"
	GetLiquorByRandomKey          = "REPO-LIQUOR-003-GetLiquorByRandomKey"
	GetLiquorByIds                = "REPO-LIQUOR-004-GetLiquorByIds"
	GetLiquorByIdsDecode          = "REPO-LIQUOR-005-GetLiquorByIdsDecode"
	LiquorCollectionCount         = "REPO-LIQUOR-006-LiquorCollectionCount"
	GetLiquorsRandom              = "REPO-LIQUOR-007-GetLiquorsRandom"
	GetLiquorsRandomDecode        = "REPO-LIQUOR-008-GetLiquorsRandomDecode"
	GetLiquorsRandomByKey         = "REPO-LIQUOR-009-GetLiquorsRandomByKey"
	GetLiquorsRandomByKeyDecode   = "REPO-LIQUOR-010-GetLiquorsRandomByKeyDecode"
	GetLiquorsRandomByKeyLt       = "REPO-LIQUOR-011-GetLiquorsRandomByKeyLt"
	GetLiquorsRandomByKeyLtDecode = "REPO-LIQUOR-012-GetLiquorsRandomByKeyLtDecode"

	GetLiquorsFromCategoryIds       = "REPO-LIQUOR-013-GetLiquorsFromCategoryIds"
	GetLiquorsFromCategoryIdsDecode = "REPO-LIQUOR-014-GetLiquorsFromCategoryIdsDecode"

	InsertOne     = "REPO-LIQUOR-015-InsertOne"
	GetInsertedId = "REPO-LIQUOR-016-GetInsertedId"

	UpdateOneBsonMap = "REPO-LIQUOR-017-UpdateOneBsonMap"
	UpdateOneToBsonM = "REPO-LIQUOR-018-UpdateOneToBsonM"
	UpdateOneExe     = "REPO-LIQUOR-019-UpdateOneExe"
	NullUpdate       = "REPO-LIQUOR-020-NullUpdate"

	DeleteRate = "REPO-LIQUOR-021-DeleteRate"
	UpdateRate = "REPO-LIQUOR-022-UpdateRate"
)

func errGetLiquorById(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorById,
		UserMsg:    "指定されたIDのお酒はありません",
		Level:      logrus.InfoLevel,
	})
}
func errGetLiquorByIds(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorByIds,
		UserMsg:    "指定されたIDのお酒はありません",
		Level:      logrus.InfoLevel,
	})
}

func errGetLiquorByName(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorByName,
		UserMsg:    "指定されたIDのお酒はありません",
		Level:      logrus.InfoLevel,
	})
}
func errGetLiquorByRandomKey(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorByRandomKey,
		UserMsg:    "指定されたIDのお酒はありません",
		Level:      logrus.InfoLevel,
	})
}
func errGetLiquorByIdsDecode(err error, ids []primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorByIdsDecode,
		UserMsg:    "デコードに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      ids,
	})
}

func errLiquorCollectionCount(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    LiquorCollectionCount,
		UserMsg:    "一覧取得に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}

func errGetLiquorsRandom(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsRandom,
		UserMsg:    "一覧取得に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}
func errGetLiquorsRandomDecode(err error, collections []*Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsRandomDecode,
		UserMsg:    "デコードに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      collections,
	})
}
func errGetLiquorsRandomByKeyDecode(err error, collections []*Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsRandomByKeyDecode,
		UserMsg:    "デコードに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      collections,
	})
}
func errGetLiquorsRandomByKey(err error, randomKey float64) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsRandomByKey,
		UserMsg:    "一覧取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      randomKey,
	})
}
func errGetLiquorsRandomByKeyLt(err error, randomKey float64) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsRandomByKeyLt,
		UserMsg:    "一覧取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      randomKey,
	})
}
func errGetLiquorsRandomByKeyLtDecode(err error, collections []*Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsRandomByKeyLtDecode,
		UserMsg:    "デコードに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      collections,
	})
}

func errGetLiquorsFromCategoryIds(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsFromCategoryIds,
		UserMsg:    "指定されたIDのお酒はありません",
		Level:      logrus.InfoLevel,
	})
}
func errGetLiquorsFromCategoryIdsDecode(err error, ids []int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetLiquorsFromCategoryIdsDecode,
		UserMsg:    "デコードに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      ids,
	})
}

func errInsertOne(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    InsertOne,
		UserMsg:    "データ追加に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}

func errGetInsertedId(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetInsertedId,
		UserMsg:    "新規追加データの取得に失敗しました",
		Level:      logrus.FatalLevel,
	})
}

func errUpdateOneBsonMap(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateOneBsonMap,
		UserMsg:    "データ更新に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}
func errUpdateOneToBsonM(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateOneToBsonM,
		UserMsg:    "データ更新に失敗しました",
		Level:      logrus.ErrorLevel,
	})
}
func errUpdateOneExe(err error, b bson.M) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateOneExe,
		UserMsg:    "データ更新に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      b,
	})
}

func errNullUpdate(b bson.M) *customError.Error {
	return customError.NewError(errors.New("no document matched the provided ID"), customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    NullUpdate,
		UserMsg:    "データ更新に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      b,
	})
}

func errDeleteRate(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    DeleteRate,
		UserMsg:    "データ更新に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}
func errUpdateRate(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateRate,
		UserMsg:    "データ更新に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}
