package logsRepository

import (
	"backend/db"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryLogRepo struct {
	db         *db.DB
	collection *mongo.Collection //コレクションを先に取得して格納しておく
}

func NewCategoryLogRepository(db *db.DB) CategoryLogRepo {
	return CategoryLogRepo{
		db:         db,
		collection: db.Collection(categoriesRepository.LogsCollectionName),
	}
}

func (r *CategoryLogRepo) GetLogsById(ctx context.Context, id int) (*[]liquorRepository.Model, error) {
	
}
