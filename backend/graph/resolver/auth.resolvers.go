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
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input graphModel.RegisterInput) (*graphModel.User, error) {
	if input.Password == nil {
		return nil, errors.New("パスワードは必須です")
	}
	//パスワードをハッシュする
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//ユーザー構造体の定義
	user := userRepository.Model{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		ImageBase64: input.ImageBase64,
		Profile:     input.Profile,
	}

	//登録して、挿入したデータを受け取る
	newUser, err := r.UserRepo.Register(ctx, &user)
	if err != nil {
		return nil, errors.New("ユーザー登録に失敗しました。")
	}
	return newUser.ToGraphQL(), nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input graphModel.RegisterInput) (bool, error) {
	loginUser, err := userService.GetUserData(ctx, r.UserRepo) //未ログイン状態ならuserIDはnilになる
	if err != nil {
		return false, err
	}
	id := loginUser.ID
	oldUser, err := r.UserRepo.GetById(ctx, id)
	if err != nil {
		return false, err
	}

	//新しいパスワードを生成する(入力が空であれば前の値を代入する)
	var newPassword []byte

	if input.Password != nil && len(*input.Password) != 0 { //空文字もnilと同等に扱う
		if len(*input.Password) < 8 {
			return false, errors.New("パスワードが短いです")
		}
		//パスワードをハッシュする
		newPassword, err = bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			return false, err
		}
	} else {
		newPassword = oldUser.Password
	}
	//ユーザー構造体の定義
	user := &userRepository.Model{
		ID:          oldUser.ID,
		Name:        input.Name,
		Email:       input.Email,
		Password:    newPassword,
		ImageBase64: input.ImageBase64,
		Profile:     input.Profile,
	}

	err = r.UserRepo.Update(ctx, user)
	if err != nil {
		return false, nil
	}

	return true, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input graphModel.LoginInput) (*graphModel.AuthPayload, error) {
	user, err := userService.Login(ctx, input, &r.UserRepo)
	if err != nil {
		return nil, err
	}
	result := &graphModel.AuthPayload{
		User:  user.User.ToGraphQL(),
		Token: user.Token,
	}
	return result, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*graphModel.User, error) {
	userID, err := userService.GetUserId(ctx)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	// ユーザー情報をデータベースから取得する処理
	user, err := r.UserRepo.GetById(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return user.ToGraphQL(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
