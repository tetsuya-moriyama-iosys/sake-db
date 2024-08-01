package resolver

import (
    "go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
    //Collection *mongo.Collection
    DB *mongo.Database
}