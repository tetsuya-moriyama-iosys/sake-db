package logger

import (
	"backend/customError"
	"backend/db/repository/errorRepository"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

// ロガーを初期化
var logger = logrus.New()

// DB 接続インスタンス
var repo errorRepository.ErrorsRepository

func Init(r errorRepository.ErrorsRepository) {
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
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{}) // JSON形式でログ出力
	logger.SetLevel(logrus.InfoLevel)            // エラーレベルのみ記録
}

// LogError はエラーをログに記録する
func LogError(ctx context.Context, err *customError.Error) {
	l := logger.WithFields(logrus.Fields{
		"error_id":   err.ID,
		"error_code": err.ErrorCode,
		"status":     err.StatusCode,
		"message":    err.LogMessage,
		"location":   err.Location,
		"timestamp":  err.Timestamp,
	})

	// エラーレベルに応じた処理
	if err.StatusCode >= 500 { // 深刻なエラー（500以上）はDBにも保存
		_ = writeDB(ctx, err)
		l.Error("Critical error occurred")
	} else {
		l.Warn("non-Critical error occurred")
	}
}

func writeDB(ctx context.Context, err *customError.Error) error {
	return repo.Write(ctx, &errorRepository.Model{
		ID:        primitive.NewObjectID(),
		Message:   err.LogMessage,
		Location:  err.Location,
		CreatedAt: time.Now(),
	})
}
