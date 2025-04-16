package categoriesRepository

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	Find            = "REPO-CATEGORY-001-Find"
	FindDecode      = "REPO-CATEGORY-002-FindDecode"
	FindCursor      = "REPO-CATEGORY-003-FindCursor"
	FindById        = "REPO-CATEGORY-004-FindById"
	NotFound        = "REPO-CATEGORY-005-NotFound"
	InsertOne       = "REPO-CATEGORY-006-InsertOne"
	InsertOneGetId  = "REPO-CATEGORY-007-InsertOneGetId"
	UpdateBsonMap   = "REPO-CATEGORY-008-UpdateBsonMap"
	UpdateUnMarshal = "REPO-CATEGORY-009-UpdateUnMarshal"
	UpdateOne       = "REPO-CATEGORY-010-UpdateOne"
	UpdateOneGetId  = "REPO-CATEGORY-011-UpdateOneGetId"
	GetMaxID  = "REPO-CATEGORY-012-GetMaxID"
)

func errFind(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Find,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.ErrorLevel,
	})
}

func errFindDecode(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FindDecode,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errFindCursor(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FindCursor,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errNotFound(err error, id int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    NotFound,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}
func errFindById(err error, id int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FindById,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errInsertOne(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    InsertOne,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}
func errInsertOneGetId(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    InsertOneGetId,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}
func errUpdateBsonMap(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateBsonMap,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}
func errUpdateUnMarshal(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateUnMarshal,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}

func errUpdateOne(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateOne,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}
func errUpdateOneGetId(err error, c *Model) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    UpdateOneGetId,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      c,
	})
}

func errGetMaxID(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetMaxID,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}
