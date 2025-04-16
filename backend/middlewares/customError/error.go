package customError

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

type ErrorLevel int

type Error struct {
	ID          string       `json:"id"`          // 一意のエラーID
	Level       logrus.Level `json:"level"`       // エラーレベル
	ErrorCode   string       `json:"error_code"`  // 一意のエラーコード
	StatusCode  int          `json:"status_code"` // HTTPステータスコード
	UserMessage string       `json:"message"`     // ユーザー向けメッセージ
	Location    string       `json:"location"`    // エラー発生場所
	Input       interface{}  `json:"input"`       // エラー発生時の入力値
	Timestamp   string       `json:"timestamp"`   // エラー発生時刻
	RawErr      error
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.ID, e.UserMessage)
}
func (e *Error) Unwrap() error {
	return e.RawErr
}

type Params struct {
	StatusCode  int
	ErrCode     string
	UserMsg     string
	ParentStack int // optional
	Input       interface{}
	Level       logrus.Level
}

func NewError(err error, params Params) *Error {
	skip := params.ParentStack + 3

	return &Error{
		ID:          uuid.New().String(),
		Level:       params.Level,
		StatusCode:  params.StatusCode,
		ErrorCode:   params.ErrCode,
		UserMessage: params.UserMsg,
		RawErr:      err,
		Location:    getErrorLocation(skip),
		Input:       params.Input,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
}

func getErrorLocation(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown location"
	}
	fn := runtime.FuncForPC(pc)
	funcName := "unknown function"
	if fn != nil {
		funcName = fn.Name()
	}
	return fmt.Sprintf("%s:%d [%s]", file, line, funcName)
}
