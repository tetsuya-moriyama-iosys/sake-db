package authService

import (
	"backend/db/repository/userRepository"
	"backend/middlewares/customError"
	"backend/util/amazon/ses"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// GeneratePasswordResetToken トークンを生成し、DBに格納する
func GeneratePasswordResetToken(ctx context.Context, r userRepository.UsersRepository, email string) (string, *customError.Error) {
	ran := rand.New(rand.NewSource(time.Now().UnixNano())) // 生成器を生成
	// ランダムな32バイトのスライスを作成
	tokenBytes := make([]byte, 32)

	// 生成器からバイトをランダムに埋める
	for i := range tokenBytes {
		tokenBytes[i] = byte(ran.Intn(256)) // 0~255の範囲でバイトを生成
	}

	//stringに変換する
	token := fmt.Sprintf("%x", tokenBytes)

	//DBにトークンを格納する
	err := r.SetPasswordToken(ctx, email, token)
	if err != nil {
		return "", err
	}

	// base64でエンコードしてトークンを文字列に変換
	return token, nil
}

func PasswordResetExe(ctx context.Context, r userRepository.UsersRepository, token string, password string) (*userRepository.Model, *customError.Error) {
	user, err := r.GetByPasswordToken(ctx, token)
	if err != nil {
		return nil, err
	}
	//パスワードをハッシュする
	var newPassword []byte
	newPassword, rawErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if rawErr != nil {
		return nil, errGenerateFromPassword(rawErr)
	}
	//パスワードリセットを実行する
	err = r.PasswordReset(ctx, *user, newPassword)
	return user, err
}

func ResetEmail(ctx context.Context, r userRepository.UsersRepository, email string) (bool, *customError.Error) {
	//トークンを生成しDBに格納する
	token, cErr := GeneratePasswordResetToken(ctx, r, email)
	if cErr != nil {
		return false, cErr
	}

	//生成したトークンからメールを作り送信する
	err := ses.SendPasswordReset(ctx, email, token)
	if err != nil {
		return false, errSendPasswordReset(err)
	}
	return true, nil
}
