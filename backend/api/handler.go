package api

import (
	"backend/db/repository/userRepository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	DB       *mongo.Database
	UserRepo userRepository.UsersRepository
}

// NewUserHandler 新しいLiquorHandlerを作成するコンストラクタ
func NewUserHandler(db *mongo.Database, userRepo userRepository.UsersRepository) *UserHandler {
	return &UserHandler{
		DB:       db,
		UserRepo: userRepo,
	}
}
