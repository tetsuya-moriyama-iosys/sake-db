package customError

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
	"time"
)

type Error struct {
	ID          string `json:"id"`          // 一意のエラーID
	ErrorCode   string `json:"error_code"`  // 一意のエラーコード
	StatusCode  int    `json:"status_code"` // HTTPステータスコード
	UserMessage string `json:"message"`     // ユーザー向けメッセージ
	LogMessage  string `json:"log_message"` // 内部ログ用メッセージ
	Location    string `json:"location"`    // エラー発生場所
	Timestamp   string `json:"timestamp"`   // エラー発生時刻
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s (at %s)", e.ID, e.LogMessage, e.Location)
}

func NewError(statusCode int, userMsg, logMsg string) *Error {
	return &Error{
		ID:          uuid.New().String(),
		StatusCode:  statusCode,
		UserMessage: userMsg,
		LogMessage:  logMsg,
		Location:    getErrorLocation(2),
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
