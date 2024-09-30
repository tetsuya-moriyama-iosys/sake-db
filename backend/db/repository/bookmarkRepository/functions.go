package bookmarkRepository

import (
	"backend/db/repository/agg"
	"backend/db/repository/userRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Target string

const (
	UserId       Target = UserID
	BookmarkedId Target = BookmarkedUserId
)

func generatePipeline(uid primitive.ObjectID, target Target, joinTarget Target) bson.A {
	return bson.A{
		// ドキュメントをフィルタリング
		agg.Where(string(target), uid),
		agg.LookUp(userRepository.CollectionName, string(joinTarget), userRepository.ID, "user_data"),
		agg.GetFirst("user_data", false),
		bson.M{"$sort": bson.M{
			"_id": -1, // _idで降順ソート（新しい順）
		}},
		bson.M{"$addFields": bson.M{
			CreatedAT: bson.M{"$toDate": "$$ROOT._id"}, // _id からタイムスタンプを生成し、created_atフィールドに追加
		}},
		//projectで整形する
		bson.M{"$project": bson.M{
			UserID:    "$user_data." + userRepository.ID, // usersコレクションからのuser_name
			UserName:  "$user_data." + userRepository.NAME,
			CreatedAT: 1,
		}},
	}
}
