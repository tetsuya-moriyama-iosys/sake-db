package helper

import (
	"github.com/go-playground/validator/v10"
)

// グローバルなバリデーションインスタンス
var validate *validator.Validate

func init() {
	// バリデーションインスタンスの初期化
	validate = validator.New()
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
