package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var youtubeURLRegex = regexp.MustCompile(`^(https?://)?(www\.)?(youtube\.com|youtu\.be)/(watch\?v=|embed/|v/|.+\?v=)?([^&=%\?]{11})$`)

// YoutubeURL バリデーション
func YoutubeURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	return youtubeURLRegex.MatchString(url)
}
