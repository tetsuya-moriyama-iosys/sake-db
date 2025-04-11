package flavorMapRepository

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	MasterFind           = "REPO-FLAVOR-MAP-001-MasterFind"
	MasterFindCursor     = "REPO-FLAVOR-MAP-002-MasterFindCursor"
	Insert               = "REPO-FLAVOR-MAP-003-Insert"
	Update               = "REPO-FLAVOR-MAP-004-Update"
	GetVotedDataByLiquor = "REPO-FLAVOR-MAP-005-GetVotedDataByLiquor"
	Upsert               = "REPO-FLAVOR-MAP-006-Upsert"
)

func errMasterFind(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    MasterFind,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.ErrorLevel,
	})
}

func errMasterFindCursor(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    MasterFindCursor,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errInsert(err error, d FlavorMapModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Insert,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      d,
	})
}

func errUpdate(err error, d FlavorMapModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Update,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      d,
	})
}

func errGetVotedDataByLiquor(err error, uId primitive.ObjectID, lId primitive.ObjectID, cId int) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetVotedDataByLiquor,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("uId: %s, lId: %s, cId: %d", uId.Hex(), lId.Hex(), cId),
	})
}

func errUpsert(err error, tying TyingModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    Upsert,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      tying,
	})
}
