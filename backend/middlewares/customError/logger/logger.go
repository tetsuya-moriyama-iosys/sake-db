package logger

import (
	"backend/db/repository/errorRepository"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
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
	uid := auth.GetIdNullable(ctx)
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

func writeDB(ctx context.Context, err *customError.Error) error {
	uid := auth.GetIdNullable(ctx)
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

func toSafeString(input interface{}) string {
	if input == nil {
		return ""
	}
	return fmt.Sprintf("%v", input)
}
