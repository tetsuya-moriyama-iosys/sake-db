package authService

import (
	"backend/db/repository/userRepository"
	"context"
	"fmt"
	"math/rand"
	"time"
)

// GeneratePasswordResetToken トークンを生成し、DBに格納する
func GeneratePasswordResetToken(ctx context.Context, r userRepository.UsersRepository, email string) (string, error) {
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
