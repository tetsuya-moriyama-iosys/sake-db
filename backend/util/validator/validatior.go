package validator

import (
	"github.com/go-playground/validator/v10"
)

// グローバルなバリデーションインスタンス
var validate *validator.Validate

func init() {
	// バリデーションインスタンスの初期化
	validate = validator.New()

	// カスタムバリデーションの登録
	err := validate.RegisterValidation("youtube", YoutubeURL)
	if err != nil {
		return
	}
}

// Validate バリデーションのヘルパー関数
func Validate(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		// バリデーションエラーがある場合、ValidationErrorsを返す
		return err.(validator.ValidationErrors)
	}
	return nil
}
