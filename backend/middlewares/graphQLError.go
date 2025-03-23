package middlewares

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/logger"
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"time"
)

// GraphQLErrorPresenter のエラーをログに記録
func GraphQLErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	// GraphQL のエラーログを記録
	log.Printf("[GraphQL ERROR]: %v", err)

	// customError を作成
	var customErr *customError.Error
	if !errors.As(err, &customErr) {
		// 未定義のエラーが存在するので、エラーコードを設定
		customErr = &customError.Error{
			ID:          fmt.Sprintf("error-%d", time.Now().Unix()),
			ErrorCode:   "GRAPHQL_ERROR",
			StatusCode:  http.StatusInternalServerError, // 未定義エラーは500として処理する
			UserMessage: fmt.Sprintf("未定義のエラー: %v", err),
			Location:    "GraphQL Resolver",
			Timestamp:   time.Now().String(),
			RawErr:      err,
		}
	}

	// gqlerror に変換
	gqlErr := graphql.DefaultErrorPresenter(ctx, err)
	gqlErr.Message = customErr.UserMessage
	gqlErr.Extensions = map[string]interface{}{
		"code":       customErr.StatusCode,
		"statusCode": customErr.StatusCode,
		"error_id":   customErr.ErrorCode,
	}

	// クライアントへ返す GraphQL エラー
	return gqlErr
}

// GraphQLRecover `panic` をキャッチする
func GraphQLRecover(ctx context.Context, err interface{}) error {
	// エラーログ記録（重大なエラーなら DB にも保存）
	logger.LogPanic(ctx, err, "GraphQL")

	// クライアントへ返す GraphQL エラー
	return gqlerror.Errorf("Internal server error")
}
