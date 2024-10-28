package flavorMapRepository

import (
	"backend/db"
)

const (
	FlavorMapCollectionName          = "flavor_map"
	FlavorMapMasterCollectionName    = "flavor_map_master"
	FlavorMapToLiquorsCollectionNAme = "flavor_map_liquors"
)

type FlavorMapRepository struct {
	db.Base
}
type FlavorMapMasterRepository struct {
	db.Base
}
type FlavorToLiquorRepository struct {
	db.Base
}

func NewFlavorMapMasterRepository(database *db.DB) FlavorMapMasterRepository {
	return FlavorMapMasterRepository{
		Base: db.Base{
			Db:         database,
			Collection: database.Collection(FlavorMapMasterCollectionName),
		},
	}
}

func NewFlavorMapRepository(database *db.DB) FlavorMapRepository {
	return FlavorMapRepository{
		Base: db.Base{
			Db:         database,
			Collection: database.Collection(FlavorMapCollectionName),
		},
	}
}

func NewFlavorToLiquorRepository(database *db.DB) FlavorToLiquorRepository {
	return FlavorToLiquorRepository{
		Base: db.Base{
			Db:         database,
			Collection: database.Collection(FlavorMapToLiquorsCollectionNAme),
		},
	}
}
