package userService

import (
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"backend/middlewares"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserWithToken struct {
	User  *userRepository.Model
	Token string
}

var ExpireTime = 3660

func Login(ctx context.Context, input graphModel.LoginInput, r *userRepository.UsersRepository) (*UserWithToken, error) {
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

	// JWTトークン生成
	expirationTime := time.Now().Add(time.Duration(ExpireTime) * time.Minute)
	claims := &middlewares.Claims{
		Id: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middlewares.JwtKey)
	if err != nil {
		return nil, err
	}
	result := &UserWithToken{
		User:  user,
		Token: tokenString,
	}
	return result, nil
}

func (u *UserWithToken) ToGraphQL() *graphModel.AuthPayload {
	return &graphModel.AuthPayload{
		User:  u.User.ToGraphQL(),
		Token: u.Token,
	}
}
