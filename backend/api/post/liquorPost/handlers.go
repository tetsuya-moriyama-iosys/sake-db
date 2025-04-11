package liquorPost

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DB           *mongo.Database
	S3Client     *s3.S3
	CategoryRepo categoriesRepository.CategoryRepository
	LiquorsRepo  liquorRepository.LiquorsRepository
	UserRepo     userRepository.UsersRepository
}

// NewHandler 新しいLiquorHandlerを作成するコンストラクタ
func NewHandler(db *mongo.Database, s3Client *s3.S3, categoryRepo categoriesRepository.CategoryRepository, liquorsRepo liquorRepository.LiquorsRepository, userRepo userRepository.UsersRepository) *Handler {
	return &Handler{
		DB:           db,
		S3Client:     s3Client,
		CategoryRepo: categoryRepo,
		LiquorsRepo:  liquorsRepo,
		UserRepo:     userRepo,
	}
}
