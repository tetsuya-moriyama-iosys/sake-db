package authService

import (
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"backend/service/authService/tokenConfig"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserWithToken struct {
	User        *userRepository.Model
	AccessToken string
}

func generateUserWithToken(user *userRepository.Model, token string) *UserWithToken {
	return &UserWithToken{
		User:        user,
		AccessToken: token,
	}
}

func (u *UserWithToken) ToGraphQL() *graphModel.AuthPayload {
	return &graphModel.AuthPayload{
		User:        u.User.ToGraphQL(),
		AccessToken: u.AccessToken,
	}
}

func getUserByInput(ctx context.Context, input graphModel.LoginInput, r *userRepository.UsersRepository) (*userRepository.Model, error) {
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

func loginById(ctx context.Context, id primitive.ObjectID, writer *http.ResponseWriter, tokenConfig tokenConfig.TokenConfig, r *userRepository.UsersRepository) (*UserWithToken, error) {
	// ユーザーインスタンスを取得
	user, err := r.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	// JWTトークン生成
	accessToken, err := GenerateTokens(writer, user.ID, tokenConfig)
	if err != nil {
		return nil, err
	}

	//UserWithTokenを生成し返す
	return generateUserWithToken(user, *accessToken), nil
}
