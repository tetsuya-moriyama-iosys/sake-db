package di

import (
	"backend/db"
	"backend/di/handlers"
	"backend/graph"
	"backend/graph/resolver"
	"backend/router"
	"backend/util/amazon/s3"
	"github.com/google/wire"
)

// BasicSet システム根幹部分
var BasicSet = wire.NewSet(
	s3.NewS3Client,
	resolver.NewResolver,
	handlers.NewHandlers,
	router.Router,
	graph.NewGraphQLServer,
	DatabaseSet,
)

// DatabaseSet データベース根幹部分
var DatabaseSet = wire.NewSet(
	db.NewMongoClient,
	db.NewDB,
	db.ProvideMongoDatabase,
)
