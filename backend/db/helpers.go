package db

import "go.mongodb.org/mongo-driver/bson"

// StructToBsonM は任意の構造体を bson.M に変換するヘルパー関数です。
func StructToBsonM(input interface{}) (bson.M, error) {
	// 構造体を BSON にマッピング
	data, err := bson.Marshal(input)
	if err != nil {
		return nil, err
	}

	// BSON を bson.M に変換
	var result bson.M
	if err := bson.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
