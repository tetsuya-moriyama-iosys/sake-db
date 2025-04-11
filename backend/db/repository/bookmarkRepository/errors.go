package bookmarkRepository

import (
	"backend/middlewares/customError"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	FindOne                 = "REPO-BOOKMARK-001-FindOne"
	ListAggregate           = "REPO-BOOKMARK-002-ListAggregate"
	ListDecode              = "REPO-BOOKMARK-003-ListDecode"
	BookmarkedListAggregate = "REPO-BOOKMARK-004-BookmarkedListAggregate"
	BookmarkedListDecode    = "REPO-BOOKMARK-005-ListDecode"
	Duplicated              = "REPO-BOOKMARK-006-Duplicated"
	ErrorOnAddFind          = "REPO-BOOKMARK-007-ErrorOnAddFind"
	ErrorOnAdd              = "REPO-BOOKMARK-008-ErrorOnAdd"
	DeleteOne               = "REPO-BOOKMARK-009-DeleteOne"
	ErrOnDelete             = "REPO-BOOKMARK-010-ErrOnDelete"
	FindRecommend           = "REPO-BOOKMARK-011-FindRecommend"
	RecommendDecode         = "REPO-BOOKMARK-012-RecommendDecode"
	RecommendErr            = "REPO-BOOKMARK-013-Recommend"
)

func errFindOne(err error, uid primitive.ObjectID, targetId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FindOne,
		UserMsg:    "ブックマークが見つかりません",
		Level:      logrus.InfoLevel,
		Input:      fmt.Printf("{uid:%v,targetId:%v}", uid, targetId),
	})
}

func errListAggregate(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ListAggregate,
		UserMsg:    "ブックマーク取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}

func errListDecode(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ListDecode,
		UserMsg:    "ブックマーク取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}

func errBookmarkedListAggregate(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BookmarkedListAggregate,
		UserMsg:    "ブックマーク取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}

func errBookmarkedListDecode(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    BookmarkedListDecode,
		UserMsg:    "ブックマーク取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}

func errDuplicated(err error, uid primitive.ObjectID, lId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    Duplicated,
		UserMsg:    "すでにブックマークされています",
		Level:      logrus.InfoLevel,
		Input:      fmt.Printf("{uid:%v,targetId:%v}", uid, lId),
	})
}

func errOnAddFind(err error, uid primitive.ObjectID, lId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ErrorOnAddFind,
		UserMsg:    "ブックマーク追加に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("{uid:%v,targetId:%v}", uid, lId),
	})
}
func errOnAdd(err error, uid primitive.ObjectID, lId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ErrorOnAdd,
		UserMsg:    "ブックマーク追加に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("{uid:%v,targetId:%v}", uid, lId),
	})
}
func errDeleteOne(err error, uid primitive.ObjectID, lId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    DeleteOne,
		UserMsg:    "ブックマーク削除に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("{uid:%v,targetId:%v}", uid, lId),
	})
}
func errOnDelete(err error, uid primitive.ObjectID, lId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ErrOnDelete,
		UserMsg:    "ブックマーク削除に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      fmt.Printf("{uid:%v,targetId:%v}", uid, lId),
	})
}
func errFindRecommend(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FindRecommend,
		UserMsg:    "データ取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}
func errRecommendDecode(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    RecommendDecode,
		UserMsg:    "データ取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}
func errRecommend(err error, uid primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    RecommendErr,
		UserMsg:    "データ取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      uid,
	})
}
