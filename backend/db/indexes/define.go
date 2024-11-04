package indexes

import (
	"backend/db/repository/bookmarkRepository"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"go.mongodb.org/mongo-driver/bson"
)

// IndexDefinitions インデックス定義のスライス
var IndexDefinitions = []IndexDefinition{
	//memo:_idには明示的にインデックスをつける必要はない(最初からユニーク制約インデックスがついてる)
	//カテゴリ
	//memo:"その他"の扱いが難しい･･････
	//{
	//	CollectionName: categoriesRepository.CollectionName,
	//	IndexKeys:      bson.D{{categoriesRepository.Name, 1}},
	//},
	{
		CollectionName: liquorRepository.CollectionName,
		IndexKeys:      bson.D{{liquorRepository.Name, 1}},
	},
	{
		CollectionName: liquorRepository.CollectionName,
		IndexKeys:      bson.D{{liquorRepository.RandomKey, 1}},
	},

	//ブックマーク類
	{
		CollectionName: bookmarkRepository.CollectionName,
		IndexKeys:      bson.D{{bookmarkRepository.UserID, 1}, {bookmarkRepository.BookmarkedUserId, 1}},
	},
	{
		CollectionName: bookmarkRepository.CollectionName,
		IndexKeys:      bson.D{{bookmarkRepository.BookmarkedUserId, 1}}, //ブックマークユーザー側からの検索もありうる
		IsNonUnique:    true,
	},

	//ログ類(バージョン番号との複合キー)
	{
		CollectionName: categoriesRepository.LogsCollectionName,
		IndexKeys:      bson.D{{categoriesRepository.ID, 1}, {categoriesRepository.VersionNo, 1}},
	},
	{
		CollectionName: liquorRepository.LogsCollectionName,
		IndexKeys:      bson.D{{liquorRepository.LiquorID, 1}, {liquorRepository.VersionNo, 1}},
	},

	//掲示板
	{
		CollectionName: liquorRepository.BoardCollectionName,
		IndexKeys:      bson.D{{liquorRepository.LiquorID, 1}},
		IsNonUnique:    true,
		//null除外はサポートされていないらしい。メモとして残しておく。→つまり、ダブりをDB側でエラー化する術がない･･････
		//PartialFilter: bson.D{
		//	{liquorRepository.UserID, bson.D{{"$ne", nil}}}, // UserIDがnullでない場合にのみ適用
		//},
	},
	{
		CollectionName: liquorRepository.TagCollectionName,
		IndexKeys:      bson.D{{liquorRepository.LiquorID, 1}},
		IsNonUnique:    true,
	},

	//タグ
	{
		CollectionName: liquorRepository.BoardCollectionName,
		IndexKeys:      bson.D{{liquorRepository.UserID, 1}},
		IsNonUnique:    true,
	},

	//ユーザー系
	{
		CollectionName: userRepository.CollectionName,
		IndexKeys:      bson.D{{userRepository.Email, 1}},
	},
}
