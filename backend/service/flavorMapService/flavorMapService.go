package flavorMapService

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/flavorMapRepository"
	"backend/db/repository/liquorRepository"
	"backend/graph/graphModel"
	"backend/graph/schema/customModel"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"backend/service/categoryService"
	"backend/util/utilType"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostFlavorMap 実際にポストする関数
func PostFlavorMap(ctx context.Context, mstR *flavorMapRepository.FlavorMapMasterRepository, flR *flavorMapRepository.FlavorToLiquorRepository, fmR *flavorMapRepository.FlavorMapRepository, cr *categoriesRepository.CategoryRepository, lr *liquorRepository.LiquorsRepository, input graphModel.PostFlavorMap, coordinates utilType.Coordinates) *customError.Error {
	lId, rawErr := primitive.ObjectIDFromHex(input.LiquorID)
	if rawErr != nil {
		return errPostFlavorMapIdFromHex(rawErr, input.LiquorID)
	}
	uId, err := auth.GetIdNullable(ctx)
	if err != nil {
		return err
	}

	//マスタデータを取得する
	mst, err := GetFlavorMasterData(ctx, mstR, lr, cr, lId)
	if err != nil {
		return err
	}
	if mst == nil {
		return errNotFoundMstData(lId)
	}

	//フレーバーマップデータを投入する
	err = fmR.PostFlavorMap(ctx, uId, lId, mst.CategoryID, coordinates)
	if err != nil {
		return err
	}

	//統計データをリフレッシュする
	err = CalcFlavorMap(ctx, mst, flR, fmR, lId)

	return err
}

// CalcFlavorMap 統計データを更新する(マスタデータは取ってきてる前提にする)
func CalcFlavorMap(ctx context.Context, mst *flavorMapRepository.MasterModel, flR *flavorMapRepository.FlavorToLiquorRepository, fmR *flavorMapRepository.FlavorMapRepository, lId primitive.ObjectID) *customError.Error {
	cursor, err := fmR.Collection.Find(ctx, bson.M{
		flavorMapRepository.LiquorID:   lId,
		flavorMapRepository.CategoryID: mst.CategoryID,
	})
	if err != nil {
		return errNotFound(err, lId, mst.CategoryID)
	}
	defer cursor.Close(ctx)
	var models []*flavorMapRepository.FlavorMapModel
	if err = cursor.All(ctx, &models); err != nil {
		return errCursor(err)
	}

	//ここから計算開始
	//有効投票数を走査
	var guestFullAmount int
	for _, model := range models {
		if model.UserId == nil {
			guestFullAmount++
		}
	}
	userFullAmount := len(models) - guestFullAmount

	//21*21の長さの配列を準備(21で割れば済むので2次元である必要がない)
	userVotes := make(map[int]int)
	guestVotes := make(map[int]int)

	// 得票数の集計
	for _, model := range models {
		xIndex := utilType.GetCoordinateIndex(model.X)
		yIndex := utilType.GetCoordinateIndex(model.Y)
		index := xIndex*21 + yIndex // 21x21のグリッドのインデックス計算

		// ユーザーごとの集計
		if model.UserId != nil {
			userVotes[index]++
		} else {
			guestVotes[index]++
		}
	}

	var cellData [21 * 21]flavorMapRepository.FlavorCellData
	var userRate = 1 //ここの倍数を弄ると登録済ユーザーの比重を変えることができる
	fullAmountForCalc := userFullAmount*userRate + guestFullAmount
	// 得票率を FlavorCellData に反映
	for i := 0; i < 21*21; i++ {
		x, y := utilType.UndoCoordinateIndex(i)
		cellData[i] = flavorMapRepository.FlavorCellData{
			Coordinates: utilType.Coordinates{X: customModel.Coordinate(x), Y: customModel.Coordinate(y)},
			Rate:        float64(userVotes[i]*userRate+guestVotes[i]) / float64(fullAmountForCalc) * 100.0, //整数除算対策
			UserAmount:  userVotes[i],
			GuestAmount: guestVotes[i],
		}
	}

	// 結果を TyingModel に反映
	cErr := flR.UpsertData(ctx, flavorMapRepository.TyingModel{
		LiquorID:        lId,
		CategoryID:      mst.CategoryID,
		FlavorCellData:  cellData,
		UserFullAmount:  userFullAmount,
		GuestFullAmount: guestFullAmount,
	})

	return cErr
}

// GetFlavorMasterData 指定されたliquorIdが属するフレーバーマップID(カテゴリID)を取得する
func GetFlavorMasterData(ctx context.Context, mstR *flavorMapRepository.FlavorMapMasterRepository, l *liquorRepository.LiquorsRepository, c *categoriesRepository.CategoryRepository, lId primitive.ObjectID) (*flavorMapRepository.MasterModel, *customError.Error) {
	liquor, err := l.GetLiquorById(ctx, lId)
	if err != nil {
		return nil, err
	}
	trail, err := categoryService.GetCategoryTrail(ctx, liquor.CategoryID, c) //パンくずリストなので順番は保証されている
	if err != nil {
		return nil, err
	}
	//マスタデータを全件取得する
	mst, err := mstR.GetMasterData(ctx)
	if err != nil {
		return nil, err
	}
	// 検索結果を保持する変数
	var foundMst *flavorMapRepository.MasterModel
	// 逆順でループ
	for trailIndex := len(*trail) - 1; trailIndex >= 0; trailIndex-- {
		// ループを使って検索
		for mstIndex := range mst {
			if mst[mstIndex].CategoryID == (*trail)[trailIndex].ID {
				foundMst = mst[mstIndex]
				break
			}
		}
	}
	return foundMst, nil //見つからなかったらnil,nilになる
}

func GetFlavorMap(ctx context.Context, mstR *flavorMapRepository.FlavorMapMasterRepository, flR *flavorMapRepository.FlavorToLiquorRepository, l *liquorRepository.LiquorsRepository, c *categoriesRepository.CategoryRepository, lId primitive.ObjectID) (*flavorMapRepository.FlavorMapResult, *customError.Error) {
	mst, err := GetFlavorMasterData(ctx, mstR, l, c, lId)
	if err != nil {
		return nil, err
	}
	if mst == nil {
		//該当するフレーバーマップがなかった場合、nil,nilで返す
		return nil, nil
	}
	var model flavorMapRepository.TyingModel
	rawErr := flR.Collection.FindOne(ctx, bson.M{flavorMapRepository.LiquorID: lId, flavorMapRepository.CategoryID: mst.CategoryID}).Decode(&model)

	if rawErr != nil {
		if !errors.Is(rawErr, mongo.ErrNoDocuments) {
			return nil, err
		}
		//見つからなかった場合は新しく作成し、それを返す
		var cellData [21 * 21]flavorMapRepository.FlavorCellData
		for i := 0; i < 21*21; i++ {
			x, y := utilType.UndoCoordinateIndex(i)
			cellData[i] = flavorMapRepository.FlavorCellData{
				Coordinates: utilType.Coordinates{X: customModel.Coordinate(x), Y: customModel.Coordinate(y)},
			}
		}
		newModel := flavorMapRepository.TyingModel{
			CategoryID:     mst.CategoryID,
			LiquorID:       lId,
			FlavorCellData: cellData,
		}
		_, rawErr = flR.Collection.InsertOne(ctx, newModel)
		if rawErr != nil {
			return nil, errInsertOne(rawErr)
		}
		return &flavorMapRepository.FlavorMapResult{
			Master: *mst,
			Tying:  newModel,
		}, err
	}

	return &flavorMapRepository.FlavorMapResult{
		Master: *mst,
		Tying:  model,
	}, err
}
