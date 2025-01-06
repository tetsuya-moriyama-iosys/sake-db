package graph

import (
	"backend/db"
	"backend/db/repository/userRepository"
	"backend/middlewares"
	"backend/service/userService"
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
)

// 必須認証のディレクティブ
func authDirective(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	// ヘッダーからトークンを取得
	tokenString, err := middlewares.ExtractTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// トークンを検証し、ユーザーIDをcontextに保存
	ctx, err = middlewares.AuthenticateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	// 認証に成功した場合、次のリゾルバを実行
	return next(ctx)
}

// 任意認証のディレクティブ
func optionalAuthDirective(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	// ヘッダーからトークンを取得
	tokenString, err := middlewares.ExtractTokenFromHeader(ctx)

	// TODO: トークンが存在しない場合とリフレッシュトークンの場合を分ける
	// トークンが存在しない場合は、認証なしで処理を続行
	if err != nil {
		return next(ctx)
	}

	// トークンを検証し、認証が成功すればユーザーIDをcontextに保存
	ctx, err = middlewares.AuthenticateToken(ctx, tokenString)
	if err != nil {
		// トークンが無効ならエラーを返す
		return nil, err
	}

	// 認証に成功した場合、次のリゾルバを実行
	return next(ctx)
}

// 管理権限認証のディレクティブ
func adminDirective(ctx context.Context, _ interface{}, next graphql.Resolver, role *string) (interface{}, error) {
	// ヘッダーからトークンを取得
	tokenString, err := middlewares.ExtractTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// トークンを検証し、ユーザーIDをcontextに保存
	ctx, err = middlewares.AuthenticateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	//追加で、権限を持っているか確認
	client, err := db.NewMongoClient()
	if err != nil {
		return nil, err
	}
	r := userRepository.NewUsersRepository(db.NewDB(client))
	err = checkRole(ctx, &r, role)
	if err != nil {
		return nil, err
	}

	// 認証に成功した場合、次のリゾルバを実行
	return next(ctx)
}

// 認証したユーザーが権限を持っているか確認
func checkRole(ctx context.Context, r *userRepository.UsersRepository, role *string) error {
	loginUser, err := userService.GetUserData(ctx, *r)
	if err != nil {
		return err
	}

	// 手動でデリファレンスして比較
	for _, v := range loginUser.Roles {
		if v == *role {
			//指定されていたロールがあった場合
			return nil
		}
	}
	return errors.New("権限エラー")
}
