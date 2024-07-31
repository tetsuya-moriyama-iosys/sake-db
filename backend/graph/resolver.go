//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
    "go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
    Collection *mongo.Collection
}

