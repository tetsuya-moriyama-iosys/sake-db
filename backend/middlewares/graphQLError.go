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
	"runtime"
	"time"
)

// GraphQLErrorPresenter のエラーをログに記録
func GraphQLErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	// GraphQL のエラーログを記録
	log.Printf("[GraphQL ERROR]: %v", err)

	// gqlerror に変換
	gqlErr := graphql.DefaultErrorPresenter(ctx, err)

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

	// エラーログ記録（重大なエラーなら DB にも保存）
	logger.LogError(ctx, customErr)

	// クライアントへ返す GraphQL エラー
	return gqlErr
}

// GraphQLRecover `panic` をキャッチする
func GraphQLRecover(ctx context.Context, err interface{}) error {
	// `panic` が発生した場所を特定
	pc, file, line, ok := runtime.Caller(3) // 3階層上の関数を取得（リゾルバ関数を指す）
	funcName := "unknown"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	// `panic` の内容をエラーログとして記録
	log.Printf("[PANIC RECOVERED]: %v (at %s:%d in %s)", err, file, line, funcName)

	// customError を作成
	errorLog := &customError.Error{
		ID:          fmt.Sprintf("panic-%d", time.Now().Unix()),
		ErrorCode:   "INTERNAL_SERVER_ERROR(GRAPHQL)",
		StatusCode:  http.StatusInternalServerError,
		UserMessage: fmt.Sprintf("重大なエラー: %v", err),
		Location:    fmt.Sprintf("%s:%d in %s", file, line, funcName),
		Timestamp:   time.Now().String(),
	}

	// エラーログ記録（LogError を活用）
	logger.LogError(ctx, errorLog)

	// クライアントへ返す GraphQL エラー
	return gqlerror.Errorf("Internal server error")
}
