package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/db"
	"backend/db/repository/liquorRepository"
	"backend/graph/graphModel"
	"backend/service/categoryService"
	"backend/service/userService"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostBoard is the resolver for the postBoard field.
func (r *mutationResolver) PostBoard(ctx context.Context, input graphModel.BoardInput) (bool, error) {
	var userID *primitive.ObjectID
	var userName *string
	user, err := userService.GetUserData(ctx, r.UserRepo) //未ログイン状態ならuserIDはnilになる

	if err != nil {
		return false, err
	}

	if user != nil {
		userID = &user.ID
		userName = &user.Name
	}

	liquorId, err := primitive.ObjectIDFromHex(input.LiquorID)
	if err != nil {
		return false, err
	}

	//挿入するデータを準備
	model := &liquorRepository.BoardModel{
		//ID:        primitive.NewObjectID(),
		UserId:    userID,
		UserName:  userName, //joinする想定だから使わない想定だが、一応非正規化して取っておく
		Text:      input.Text,
		Rate:      input.Rate,
		LiquorID:  liquorId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//トランザクション(返り値を返さないといけない構造になっていたので、boolを返すことにした)
	_, err = db.WithTransaction(ctx, r.DB.Client(), func(sc mongo.SessionContext) (bool, error) {
		err = r.LiquorRepo.BoardInsert(ctx, model) //掲示板を更新する(1ユーザーについて1つ)
		if err != nil {
			return false, err
		}
		//ユーザーが存在しており、かつ評価値がある場合はupdateする
		if userID != nil {
			err = r.LiquorRepo.UpdateRate(ctx, input.LiquorID, *userID, input.Rate)
			if err != nil {
				return false, err
			}
		}
		return true, nil
	})
	return true, nil
}

// Liquor is the resolver for the liquor field.
func (r *queryResolver) Liquor(ctx context.Context, id string) (*graphModel.Liquor, error) {
	liquor, err := r.LiquorRepo.GetLiquorById(ctx, id)
	if err != nil {
		return nil, err
	}
	//所属するカテゴリのリストを取得する
	trails, err := categoryService.GetCategoryTrail(ctx, liquor.CategoryID, &r.CategoryRepo)
	if err != nil {
		return nil, err
	}

	//GraphQLが期待する型に変換
	var trailQL []*graphModel.CategoryTrail
	for _, trail := range *trails {
		t := graphModel.CategoryTrail{
			ID:   trail.ID,
			Name: trail.Name,
		}
		trailQL = append(trailQL, &t)
	}

	result := liquor.ToGraphQL()
	result.CategoryTrail = trailQL
	return result, nil
}

// RandomRecommendList is the resolver for the randomRecommendList field.
func (r *queryResolver) RandomRecommendList(ctx context.Context, limit int) ([]*graphModel.Liquor, error) {
	collection, err := r.LiquorRepo.GetRandomLiquors(ctx, limit)
	if err != nil {
		return nil, err
	}
	var results []*graphModel.Liquor

	// 結果をGraphQLの構造体にマッピング（スネークケースからキャメルケースへの変換）
	for _, liquor := range collection {
		results = append(results, liquor.ToGraphQL())
	}

	return results, nil
}

// ListFromCategory is the resolver for the listFromCategory field.
func (r *queryResolver) ListFromCategory(ctx context.Context, categoryID int) (*graphModel.ListFromCategory, error) {
	ids, err := categoryService.GetBelongCategoryIdList(ctx, categoryID, &r.CategoryRepo)
	if err != nil {
		return nil, err
	}
	list, err := r.LiquorRepo.GetLiquorsFromCategoryIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	var liquors []*graphModel.Liquor
	//GraphQLスキーマに変換
	for _, liquor := range list {
		liquors = append(liquors, liquor.ToGraphQL())
	}
	//カテゴリ名を取得する
	category, err := r.CategoryRepo.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	result := &graphModel.ListFromCategory{
		CategoryName:        category.Name,
		CategoryDescription: category.Description,
		Liquors:             liquors,
	}

	return result, err
}

// LiquorHistories is the resolver for the liquorHistories field.
func (r *queryResolver) LiquorHistories(ctx context.Context, id string) (*graphModel.LiquorHistory, error) {
	//まず対象のカテゴリ情報を取得
	liquor, err := r.LiquorRepo.GetLiquorById(ctx, id)
	if err != nil {
		return nil, err
	}
	logs, err := r.LiquorRepo.GetLogsById(ctx, id)
	if err != nil {
		return nil, err
	}

	var graphLogs []*graphModel.Liquor
	if logs != nil {
		for _, log := range logs {
			graphLogs = append(graphLogs, log.ToGraphQL())
		}
	}
	result := &graphModel.LiquorHistory{
		Now:       liquor.ToGraphQL(),
		Histories: graphLogs,
	}
	return result, nil
}

// Board TODO:ページネーション
func (r *queryResolver) Board(ctx context.Context, liquorID string, page *int) ([]*graphModel.BoardPost, error) {
	liquorIdObj, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, err
	}
	posts, err := r.LiquorRepo.BoardList(ctx, liquorIdObj)
	if err != nil {
		return nil, err
	}
	var result []*graphModel.BoardPost
	for _, post := range posts {
		result = append(result, post.ToGraphQL())
	}
	return result, nil
}

// GetMyBoard 自身の投稿を取得する(初期値設定用)
func (r *queryResolver) GetMyBoard(ctx context.Context, liquorID string) (*graphModel.BoardPost, error) {
	//未ログイン時にも呼ばれる関数であり、未ログインはエラーなしで空値を返すという処理をする必要がある
	isLogin := userService.IsLogin(ctx)
	if isLogin == false {
		return nil, nil
	}

	//ここからはユーザーが存在している前提の処理
	id, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, err
	}
	uid, err := userService.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	board, err := r.LiquorRepo.BoardGetByUserAndLiquor(ctx, id, *uid)
	if err != nil {
		return nil, err
	}

	return board.ToGraphQL(), nil
}
