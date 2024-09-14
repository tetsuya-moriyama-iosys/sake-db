package indexes

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"go.mongodb.org/mongo-driver/bson"
)

// IndexDefinitions インデックス定義のスライス
var IndexDefinitions = []IndexDefinition{
	{
		CollectionName: categoriesRepository.LogsCollectionName,
		IndexKeys:      bson.D{{categoriesRepository.ID, 1}, {categoriesRepository.VersionNo, 1}},
	},
	{
		CollectionName: liquorRepository.LogsCollectionName,
		IndexKeys:      bson.D{{liquorRepository.LogID, 1}, {liquorRepository.VersionNo, 1}},
	},
	{
		CollectionName: userRepository.CollectionName,
		IndexKeys:      bson.D{{userRepository.EMAIL, 1}},
	},
}
