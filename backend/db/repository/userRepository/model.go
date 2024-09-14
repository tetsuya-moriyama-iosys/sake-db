package userRepository

import (
	"backend/graph/graphModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionName = "users"
	ID             = "_id"
	NAME           = "name"
	EMAIL          = "email"
	PASSWORD       = "password"
)

type Model struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func (m *Model) ToGraphQL() *graphModel.User {
	return &graphModel.User{
		ID:    m.ID.Hex(),
		Name:  m.Name,
		Email: m.Email,
	}
}
