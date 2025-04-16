package flavorMapRepository

import (
	"backend/middlewares/customError"
	"backend/util/utilType"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// GetMasterData フルのマスターデータを取得する
func (r *FlavorMapMasterRepository) GetMasterData(ctx context.Context) ([]*MasterModel, *customError.Error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errMasterFind(err)
	}

	var models []*MasterModel
	if err = cursor.All(ctx, &models); err != nil {
		return nil, errMasterFindCursor(err)
	}

	return models, nil
}

func (r *FlavorMapRepository) PostFlavorMap(ctx context.Context, uId *primitive.ObjectID, liquorId primitive.ObjectID, categoryId int, coordinates utilType.Coordinates) *customError.Error {
	d := FlavorMapModel{
		LiquorId:   liquorId,
		CategoryId: categoryId,
		UserId:     uId,
		X:          coordinates.X,
		Y:          coordinates.Y,
		UpdatedAt:  time.Now(),
	}
	var err error
	if uId == nil {
		_, err = r.Collection.InsertOne(ctx, d)
		if err != nil {
			return errInsert(err, d)
		}
	} else {
		_, err = r.Upsert(ctx, bson.M{UserID: uId, CategoryID: categoryId, LiquorID: liquorId}, d)
		if err != nil {
			return errUpdate(err, d)
		}
	}
	return nil
}

func (r *FlavorMapRepository) GetVotedDataByLiquor(ctx context.Context, uId primitive.ObjectID, lId primitive.ObjectID, cId int) (*FlavorMapModel, *customError.Error) {
	var model *FlavorMapModel
	if err := r.Collection.FindOne(ctx, bson.M{
		UserID:     uId,
		LiquorID:   lId,
		CategoryID: cId,
	}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, errGetVotedDataByLiquor(err, uId, lId, cId)
	}

	return model, nil
}

func (r *FlavorToLiquorRepository) UpsertData(ctx context.Context, tying TyingModel) *customError.Error {
	// 結果を TyingModel に反映
	_, err := r.Upsert(ctx, bson.M{
		LiquorID:   tying.LiquorID,
		CategoryID: tying.CategoryID,
	}, tying)

	if err != nil {
		return errUpsert(err, tying)
	}

	return nil
}
