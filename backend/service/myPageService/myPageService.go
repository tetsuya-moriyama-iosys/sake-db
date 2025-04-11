package myPageService

import (
	"backend/db/repository/userRepository"
	"backend/graph/graphModel"
	"backend/middlewares/customError"
	"backend/service/userService"
	"context"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(ctx context.Context, r userRepository.UsersRepository, input graphModel.RegisterInput) (bool, *customError.Error) {
	loginUser, err := userService.GetUserData(ctx, r) //未ログイン状態ならuserIDはnilになる
	if err != nil {
		return false, err
	}
	id := loginUser.ID
	oldUser, err := r.GetById(ctx, id)
	if err != nil {
		return false, err
	}

	//新しいパスワードを生成する(入力が空であれば前の値を代入する)
	var newPassword []byte

	if input.Password != nil && len(*input.Password) != 0 { //空文字もnilと同等に扱う
		if len(*input.Password) < 8 {
			return false, errTooShortPassword()
		}
		//パスワードをハッシュする
		p, rawErr := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		newPassword = p //直接代入しようとしてもうまくいかないっぽい
		if rawErr != nil {
			return false, errGenerateFromPassword(rawErr)
		}
	} else {
		newPassword = oldUser.Password
	}
	//ユーザー構造体の定義
	user := &userRepository.Model{
		ID:          oldUser.ID,
		Name:        input.Name,
		Email:       &input.Email,
		Password:    newPassword,
		ImageBase64: input.ImageBase64,
		Profile:     input.Profile,
	}

	err = r.Update(ctx, user)
	if err != nil {
		return false, nil
	}

	return true, nil
}
