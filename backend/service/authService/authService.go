package authService

import (
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"backend/service/authService/tokenConfig"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// RefreshTokens アクセストークンが切れたため、リフレッシュトークンを使いトークンを再生成
func RefreshTokens(req *http.Request, writer http.ResponseWriter, tokenConfig tokenConfig.TokenConfig) (*string, *customError.Error) {
	return refreshHandler(req, writer, tokenConfig)
}

// LoginWithRefreshToken リフレッシュトークンを用いてログインする
func LoginWithRefreshToken(ctx context.Context, req *http.Request, writer http.ResponseWriter, tokenConfig tokenConfig.TokenConfig, r *userRepository.UsersRepository) (*UserWithToken, *customError.Error) {
	claims, err := parseRefreshToken(req, tokenConfig)
	if err != nil {
		return nil, err
	}

	// ユーザーインスタンスを取得
	return loginById(ctx, claims.Id, writer, tokenConfig, r)
}

// GenerateTokens トークンを生成
func GenerateTokens(writer http.ResponseWriter, id primitive.ObjectID, tokenConfig tokenConfig.TokenConfig) (*string, *customError.Error) {
	// アクセストークン
	accessClaims := auth.Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenConfig.AccessExpire)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(tokenConfig.AccessSecretKey)
	if err != nil {
		return nil, err
	}

	// リフレッシュトークン
	err = resetRefreshToken(writer, id, tokenConfig)
	if err != nil {
		return nil, err
	}

	return &accessString, nil
}

func DeleteRefreshToken(writer http.ResponseWriter) *customError.Error {
	//クッキーを消去
	http.SetCookie(writer, &http.Cookie{
		Name:     refreshTokenName,
		Value:    "",
		Expires:  time.Unix(0, 0), // 過去の時刻に設定
		Path:     "/",
		MaxAge:   -1, // 即座に削除
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // 必要かどうか後で確認
	})
	return nil
}

// RegisterUser is the resolver for the registerUser field.
func RegisterUser(ctx context.Context, r userRepository.UsersRepository, input graphModel.RegisterInput) (*userRepository.Model, *customError.Error) {
	//TODO: なんかロジックが大きいのでサービス層に分離すべきな気がする
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
		Email:       &input.Email,
		Password:    hashedPassword,
		ImageBase64: input.ImageBase64,
		Profile:     input.Profile,
	}

	//登録して、挿入したデータを受け取る
	newUser, err := r.Register(ctx, &user)
	if err != nil {
		// MongoDBエラーかつ重複エラーかを判定
		var mongoErr mongo.WriteException
		if errors.As(err, &mongoErr) {
			if len(mongoErr.WriteErrors) > 0 && mongoErr.WriteErrors[0].Code == 11000 {
				return nil, errors.New("このメールアドレスは既に登録されています。")
			}
		}
		return nil, errors.New("ユーザー登録に失敗しました。")
	}
	return newUser, nil
}
