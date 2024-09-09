package liquorPost

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DB           *mongo.Database
	S3Client     *s3.S3
	CategoryRepo categoriesRepository.CategoryRepository
	LiquorsRepo  liquorRepository.LiquorsRepository
}

// NewHandler 新しいLiquorHandlerを作成するコンストラクタ
func NewHandler(db *mongo.Database, s3Client *s3.S3, categoryRepo categoriesRepository.CategoryRepository, LiquorsRepo liquorRepository.LiquorsRepository) *Handler {
	return &Handler{
		DB:           db,
		S3Client:     s3Client,
		CategoryRepo: categoryRepo,
		LiquorsRepo:  LiquorsRepo,
	}
}
