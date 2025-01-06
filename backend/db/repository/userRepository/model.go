package userRepository

import (
	"backend/graph/graphModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName           = "users"
	Id                       = "_id"
	Name                     = "name"
	ImageBase64              = "image_base64"
	Email                    = "email"
	Password                 = "password"
	PasswordResetToken       = "password_reset_token"
	PasswordResetTokenExpire = "password_reset_expire"
)

type Model struct {
	ID                  primitive.ObjectID `bson:"_id"`
	Name                string             `bson:"name"`
	Email               string             `bson:"email"`
	Roles               []string           `bson:"roles"`
	Password            []byte             `bson:"password"`
	ImageBase64         *string            `bson:"image_base64"`
	Profile             *string            `bson:"profile"`
	PasswordResetToken  *[]byte            `bson:"password_reset_token"`
	PasswordResetExpire *time.Time         `bson:"password_reset_expire"`
}

func (m *Model) ToGraphQL() *graphModel.User {
	return &graphModel.User{
		ID:          m.ID.Hex(),
		Name:        m.Name,
		Email:       m.Email,
		ImageBase64: m.ImageBase64,
		Profile:     m.Profile,
		Roles:       m.Roles,
	}
}
