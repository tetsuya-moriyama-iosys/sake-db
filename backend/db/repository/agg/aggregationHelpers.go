// Package agg
/**
  agg Aggregationを使いやすいようにヘルパ化したもの
*/
package agg

import "go.mongodb.org/mongo-driver/bson"

func Where(key string, value any) bson.M {
	return bson.M{"$match": bson.M{key: value}}
}

func LookUp(from string, localField string, foreignField string, as string) bson.M {
	return bson.M{"$lookup": bson.M{
		"from":         from,         // 参照するコレクション
		"localField":   localField,   // フィールド
		"foreignField": foreignField, // 結合相手のフィールド
		"as":           as,           // 結果を格納するフィールド名
	}}
}

// GetFirst LookUpOneとしたかったが、[]bson.M...やbson.Aがbson.A内で使えなかったため仕方なく分離。指定したフィールドの0番キーを取り出す。
func GetFirst(as string) bson.M {
	return bson.M{"$unwind": bson.M{"path": "$" + as, "preserveNullAndEmptyArrays": true}}
}
