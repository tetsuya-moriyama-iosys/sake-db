package flavorMapRepository

import (
	"backend/util/utilType"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// GetMasterData フルのマスターデータを取得する
func (r *FlavorMapMasterRepository) GetMasterData(ctx context.Context) ([]*MasterModel, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var models []*MasterModel
	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	return models, nil
}

func (r *FlavorMapRepository) PostFlavorMap(ctx context.Context, uId *primitive.ObjectID, liquorId primitive.ObjectID, categoryId int, coordinates utilType.Coordinates) error {
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
	} else {
		_, err = r.Upsert(ctx, bson.M{UserID: uId, CategoryID: categoryId, LiquorID: liquorId}, d)
	}
	return err
}

func (r *FlavorMapRepository) GetVotedDataByLiquor(ctx context.Context, uId primitive.ObjectID, lId primitive.ObjectID, cId int) (*FlavorMapModel, error) {
	var model *FlavorMapModel
	if err := r.Collection.FindOne(ctx, bson.M{
		UserID:     uId,
		LiquorID:   lId,
		CategoryID: cId,
	}).Decode(&model); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return model, nil
}
