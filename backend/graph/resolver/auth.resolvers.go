package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/db/repository/userRepository"
	"backend/graph/generated"
	"backend/graph/graphModel"
	"backend/service/userService"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input graphModel.RegisterInput) (*graphModel.User, error) {
	//パスワードをハッシュする
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//ユーザー構造体の定義
	user := userRepository.Model{
		ID:       primitive.NewObjectID(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	//登録して、挿入したデータを受け取る
	newUser, err := r.UserRepo.Register(ctx, &user)
	if err != nil {
		return nil, err
	}
	return newUser.ToGraphQL(), nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input graphModel.LoginInput) (*graphModel.User, error) {
	//panic(fmt.Errorf("not implemented: Login - login"))
	user, err := userService.Login(ctx, input, &r.UserRepo)
	if err != nil {
		return nil, err
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
