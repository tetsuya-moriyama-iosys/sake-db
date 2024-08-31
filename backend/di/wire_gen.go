// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"backend/api/post/liquorPost"
	"backend/db"
	"backend/db/categoriesRepository"
	"backend/db/liquorRepository"
	"backend/di/handlers"
	"backend/graph"
	"backend/graph/resolver"
	"backend/router"
	"backend/util/amazon/s3"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitializeHandler() (*gin.Engine, error) {
	client, err := db.NewMongoClient()
	if err != nil {
		return nil, err
	}
	dbDB := db.NewDB(client)
	categoryRepository := categoriesRepository.NewCategoryRepository(dbDB)
	liquorsRepository := liquorRepository.NewLiquorsRepository(dbDB)
	resolverResolver := resolver.NewResolver(categoryRepository, liquorsRepository)
	server := graph.NewGraphQLServer(resolverResolver)
	database := db.ProvideMongoDatabase(dbDB)
	s3S3, err := s3.NewS3Client()
	if err != nil {
		return nil, err
	}
	handler := liquorPost.NewHandler(database, s3S3, categoryRepository, liquorsRepository)
	handlersHandlers := handlers.NewHandlers(handler)
	engine := router.Router(server, handlersHandlers)
	return engine, nil
}