package flavorMapRepository

import (
	"backend/graph/graphModel"
	"backend/graph/schema/customModel"
	"backend/util/utilType"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// MasterModel フレーバーマップのマスタモデル
type MasterModel struct {
	CategoryID int       `bson:"category_id"`
	XName      [2]string `bson:"x"`
	YName      [2]string `bson:"y"`
}

// TyingModel 特定のお酒とフレーバーマップを関連付けるコレクション(カテゴリが移動するとフレーバーマップが変わる可能性があるが、戻した時に復元できないといけない)
type TyingModel struct {
	LiquorID        primitive.ObjectID      `bson:"liquor_id"`
	CategoryID      int                     `bson:"category_id"`
	FlavorCellData  [21 * 21]FlavorCellData `bson:"flavor_cell_data"`
	UserFullAmount  int                     `bson:"user_full_amount"`
	GuestFullAmount int                     `bson:"guest_full_amount"`
}
type FlavorCellData struct {
	utilType.Coordinates
	Rate        float64 `bson:"rate"`
	UserAmount  int     `bson:"user_amount"`
	GuestAmount int     `bson:"guest_amount"`
}

type FlavorMapModel struct {
	LiquorId   primitive.ObjectID     `bson:"liquor_id"`
	CategoryId int                    `bson:"category_id"`
	UserId     *primitive.ObjectID    `bson:"user_id"`
	X          customModel.Coordinate `bson:"x"`
	Y          customModel.Coordinate `bson:"y"`
	UpdatedAt  time.Time              `bson:"updated_at"`
}

// FlavorMapResult フロントに渡すためにマスタデータとTyingデータの両方が必要なのでまとめる
type FlavorMapResult struct {
	Master MasterModel
	Tying  TyingModel
}

// ToGraphQL 他のリポジトリと異なり、モデルを直接GraphQLに変換するのではなく、マスタデータとフレーバーマップモデルを組み合わせてフロントが欲しい形に整形する
func (r *FlavorMapResult) ToGraphQL() *graphModel.FlavorMapData {
	var fMap []*graphModel.FlavorCellData
	for _, f := range r.Tying.FlavorCellData {
		fMap = append(fMap, f.ToGraphQL())
	}

	return &graphModel.FlavorMapData{
		//マスタデータ
		CategoryID: r.Master.CategoryID,
		XNames:     r.Master.XName[:], //gqlgenは固定長の配列に対応していないので、スライスに変換
		YNames:     r.Master.YName[:],
		//ソフトデータ
		UserFullAmount:  r.Tying.UserFullAmount,
		GuestFullAmount: r.Tying.GuestFullAmount,
		MapData:         fMap,
	}
}

func (f *FlavorCellData) ToGraphQL() *graphModel.FlavorCellData {
	return &graphModel.FlavorCellData{
		X:           f.X,
		Y:           f.Y,
		Rate:        f.Rate,
		GuestAmount: f.GuestAmount,
		UserAmount:  f.UserAmount,
	}
}

func (f *FlavorMapModel) ToGraphQL() *graphModel.VotedData {
	return &graphModel.VotedData{
		LiquorID:   f.LiquorId.Hex(),
		UserID:     f.UserId.Hex(),
		CategoryID: f.CategoryId,
		X:          f.X,
		Y:          f.Y,
		UpdatedAt:  f.UpdatedAt,
	}
}
