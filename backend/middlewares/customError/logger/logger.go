package logger

import (
	"backend/db/repository/errorRepository"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

// ロガーを初期化
var logger = logrus.New()

// DB 接続インスタンス
var repo errorRepository.ErrorsRepository

func Init(r errorRepository.ErrorsRepository) {
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic(fmt.Sprintf("ログディレクトリ作成に失敗: %v", err))
	}
	repo = r

	// `lumberjack` を設定
	logFile := &lumberjack.Logger{
		Filename:   "logs/err.log", // ログファイルのパス
		MaxSize:    10,             // 10MB以上になったらローテーション
		MaxBackups: 30,             // 最大30個のログファイルを保存
		MaxAge:     7,              // 7日以上のログは削除
		Compress:   false,          // 圧縮するかどうか
	}

	// logrus の出力をファイルに変更
	logger.SetOutput(logFile)
	logger.SetFormatter(&logrus.JSONFormatter{}) // JSON形式でログ出力
	logger.SetLevel(logrus.WarnLevel)
}

// LogError はエラーをログに記録する
func LogError(ctx context.Context, err *customError.Error) {
	uid, _ := auth.GetIdNullable(ctx) //ID取得エラーは一旦握りつぶす。問題があればフィールドに追加して対応する。

	l := logger.WithFields(logrus.Fields{
		"error_id":   err.ID,
		"error_code": err.ErrorCode,
		"user_id":    toSafeString(uid),
		"status":     err.StatusCode,
		"message":    err.RawErr.Error(),
		"input":      toSafeString(err.Input),
		"location":   err.Location,
		"timestamp":  err.Timestamp,
	})
	fmt.Printf("エラー: %v\n", l)

	// エラーレベルに応じた処理
	go func() {
		ctxTO, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // タイムアウト後に `ctx` を解放

		if err.Level <= logrus.ErrorLevel { // 深刻なエラーはDBにも保存
			_ = writeDB(ctxTO, err)
			l.Error("Critical error occurred")
		} else {
			l.Warn("non-Critical error occurred")
		}
	}()
}

// LogPanic 共通のpanic処理
func LogPanic(ctx context.Context, recovered interface{}, locationHint string) customError.Error {
	pc, file, line, ok := runtime.Caller(3)
	funcName := "unknown"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	stack := string(debug.Stack())
	id := primitive.NewObjectID().Hex()
	errCode := fmt.Sprintf("panic-%s", primitive.NewObjectID().Hex())

	customErr := customError.Error{
		ID:          id,
		ErrorCode:   errCode,
		StatusCode:  500,
		Level:       logrus.ErrorLevel,
		UserMessage: fmt.Sprintf("内部エラーが発生しました。 [エラーコード：%s]", errCode),
		RawErr:      errors.New(fmt.Sprintf("%v\n%s", recovered, stack)),
		Location:    fmt.Sprintf("%s:%d in %s [%s]", file, line, funcName, locationHint),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	LogError(ctx, &customErr)
	return customErr
}

func writeDB(ctx context.Context, err *customError.Error) error {
	uid, _ := auth.GetIdNullable(ctx)
	return repo.Write(ctx, &errorRepository.Model{
		ID:        primitive.NewObjectID(),
		Code:      err.ErrorCode,
		UserId:    uid,
		Message:   err.RawErr.Error(),
		Location:  err.Location,
		Input:     toSafeString(err.Input),
		CreatedAt: time.Now(),
	})
}

const maxStringLength = 1000

func toSafeString(input interface{}) (s string) {
	defer func() {
		if r := recover(); r != nil {
			s = "[error converting to string]"
		}
	}()

	if input == nil {
		return ""
	}

	switch v := input.(type) {
	case fmt.Stringer:
		s = v.String()
	case string:
		s = v
	case []byte:
		if len(v) > 1024 {
			s = fmt.Sprintf("[byte slice too long: %d bytes]", len(v))
		} else {
			s = string(v)
		}
	case error:
		s = v.Error()
	default:
		s = fmt.Sprintf("%#v", v)
	}

	// 文字列長を制限
	if len(s) > maxStringLength {
		s = s[:maxStringLength] + "...[truncated]"
	}

	return s
}
