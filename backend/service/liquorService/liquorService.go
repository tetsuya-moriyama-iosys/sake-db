package liquorService

import (
	"backend/db"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"backend/middlewares/customError"
	"backend/service/categoryService"
	"backend/service/userService"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetLiquor(ctx context.Context, lr liquorRepository.LiquorsRepository, cr categoriesRepository.CategoryRepository, id string) (*graphModel.Liquor, *customError.Error) {
	lid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errGetLiquorIdHex(err, id)
	}
	liquor, err := lr.GetLiquorById(ctx, lid)
	if err != nil {
		return nil, errGetLiquorId(err, lid)
	}
	//所属するカテゴリのリストを取得する
	trails, cErr := categoryService.GetCategoryTrail(ctx, liquor.CategoryID, &cr)
	if cErr != nil {
		return nil, cErr
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

func PostBoard(ctx context.Context, lr liquorRepository.LiquorsRepository, ur userRepository.UsersRepository, input graphModel.BoardInput) (bool, *customError.Error) {
	//バリデーション処理
	if len(input.Text) > 500 {
		return false, nil
	}
	if input.Rate != nil && (*input.Rate < 1 || *input.Rate > 5) {
		return false, nil
	}

	var userID *primitive.ObjectID                //名無しの可能性がある
	user, err := userService.GetUserData(ctx, ur) //未ログイン状態ならuserIDはnilになる

	if err != nil {
		return false, err
	}

	if user != nil {
		userID = &user.ID
	}

	lId, e := primitive.ObjectIDFromHex(input.LiquorID)
	if e != nil {
		return false, errPostBoardObjectIDFromHex(e, input.LiquorID)
	}

	//挿入するデータを準備
	model := &liquorRepository.BoardModel{
		UserId:    userID,
		LiquorID:  lId,
		Text:      input.Text,
		Rate:      input.Rate,
		UpdatedAt: time.Now(),
	}

	//トランザクション(返り値を返さないといけない構造になっていたので、boolを返すことにした)
	_, e = db.WithTransaction(ctx, lr.DB.Client, func(sc mongo.SessionContext) (bool, error) {
		err = lr.BoardInsert(ctx, model) //掲示板を更新する(1ユーザーについて1つ)
		if err != nil {
			return false, err
		}
		//ユーザーが存在しており、かつ評価値がある場合はupdateする
		if userID != nil {
			err = lr.UpdateRate(ctx, lId, *userID, input.Rate)
			if err != nil {
				return false, err
			}
		}
		return true, nil
	})
	if e != nil {
		return false, errPostBoard(e, model)
	}
	return true, nil
}

func GetLiquorHistories(ctx context.Context, r liquorRepository.LiquorsRepository, id string) (*graphModel.LiquorHistory, *customError.Error) {
	lid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errGetLiquorHistoriesIDFromHex(err, id)
	}
	//まず対象のカテゴリ情報を取得
	liquor, cErr := r.GetLiquorById(ctx, lid)
	if cErr != nil {
		return nil, cErr
	}
	logs, cErr := r.GetLogsById(ctx, lid)
	if cErr != nil {
		return nil, cErr
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

// GetBoard TODO:ページネーション
func GetBoard(ctx context.Context, r liquorRepository.LiquorsRepository, liquorID string, page *int) ([]*graphModel.BoardPost, *customError.Error) {
	liquorIdObj, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, errGetBoardFromHex(err, liquorID)
	}
	posts, cErr := r.BoardList(ctx, liquorIdObj)
	if cErr != nil {
		return nil, cErr
	}
	var result []*graphModel.BoardPost
	for _, post := range posts {
		result = append(result, post.ToGraphQL())
	}
	return result, nil
}

// GetMyBoard 自身の投稿を取得する(初期値設定用)
func GetMyBoard(ctx context.Context, r liquorRepository.LiquorsRepository, liquorID string, uId primitive.ObjectID) (*liquorRepository.BoardModel, *customError.Error) {
	id, err := primitive.ObjectIDFromHex(liquorID)
	if err != nil {
		return nil, errGetMyBoard(err, liquorID)
	}

	board, rErr := r.BoardGetByUserAndLiquor(ctx, id, uId)
	if rErr != nil {
		// 結果が0件の場合、nilを返す
		if errors.Is(rErr.RawErr, mongo.ErrNoDocuments) {
			return nil, nil
		}
		// 他のエラーの場合はそのまま返す
		return nil, rErr
	}

	//対象が存在しなければ、普通にnilを返す
	if board == nil {
		return nil, nil
	}

	return board, nil
}
