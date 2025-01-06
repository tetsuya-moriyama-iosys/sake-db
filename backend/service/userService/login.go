package userService

import (
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserWithToken struct {
	User        *userRepository.Model
	AccessToken string
}

func Login(ctx context.Context, input graphModel.LoginInput, r *userRepository.UsersRepository) (*UserWithToken, error) {
	// ユーザーインスタンスを取得
	user, err := authCheck(ctx, input, r)
	if err != nil {
		return nil, err
	}

	// JWTトークン生成
	accessToken, err := generateTokens(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	result := &UserWithToken{
		User:        user,
		AccessToken: *accessToken,
	}
	return result, nil
}

func authCheck(ctx context.Context, input graphModel.LoginInput, r *userRepository.UsersRepository) (*userRepository.Model, error) {
	// ユーザーインスタンスを取得
	user, err := r.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("メールアドレスもしくはパスワードが間違っています。")
	}

	// パスワード検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("メールアドレスもしくはパスワードが間違っています。")
	}
	return user, nil
}

func (u *UserWithToken) ToGraphQL() *graphModel.AuthPayload {
	return &graphModel.AuthPayload{
		User:        u.User.ToGraphQL(),
		AccessToken: u.AccessToken,
	}
}
