package utilType

import "backend/graph/schema/customModel"

type Coordinates struct {
	X customModel.Coordinate `bson:"x"`
	Y customModel.Coordinate `bson:"y"`
}

// GetCoordinateIndex 座標をインデックスに変換する関数 (-10～10の範囲を0～20に変換)
func GetCoordinateIndex(coordinate customModel.Coordinate) int {
	return int(coordinate + 10)
}

func UndoCoordinateIndex(index int) (int, int) {
	return index/21 - 10, index%21 - 10
}
