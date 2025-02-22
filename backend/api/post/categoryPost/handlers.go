package categoryPost

import (
	"backend/db/repository/categoriesRepository"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DB           *mongo.Database
	S3Client     *s3.S3
	CategoryRepo categoriesRepository.CategoryRepository
}

// NewHandler 新しいLiquorHandlerを作成するコンストラクタ
func NewHandler(db *mongo.Database, s3Client *s3.S3, categoryRepo categoriesRepository.CategoryRepository) *Handler {
	return &Handler{
		DB:           db,
		S3Client:     s3Client,
		CategoryRepo: categoryRepo,
	}
}
