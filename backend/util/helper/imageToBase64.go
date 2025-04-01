package helper

import (
	"backend/middlewares/customError"
	"bytes"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"image"
	"image/jpeg"
	"net/http"
)

type Base64Option struct {
	MaxWidth  *uint // リサイズ時の最大幅
	MaxHeight *uint // リサイズ時の高さ
}

func ImageToBase64(img image.Image, option *Base64Option) (*string, *customError.Error) {
	// リサイズ実行
	thumbnail := ResizeImage(img, option.MaxWidth, option.MaxHeight)

	// Base64エンコード
	var thumbBuf bytes.Buffer
	err := jpeg.Encode(&thumbBuf, thumbnail, nil)
	if err != nil {
		return nil, errImageToBase64(err, img)
	}

	//string型を*stringに変換する
	encoded := base64.StdEncoding.EncodeToString(thumbBuf.Bytes())
	imageBase64 := &encoded

	return imageBase64, nil
}

func GenerateBase64Option(h int, w int) *Base64Option {
	height := uint(h)
	width := uint(w)
	option := Base64Option{
		MaxHeight: &height,
		MaxWidth:  &width,
	}
	return &option
}

func errImageToBase64(err error, img image.Image) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode:  http.StatusBadRequest,
		ErrCode:     "ERR-ImageToBase64",
		UserMsg:     "ファイルのリサイズに失敗しました",
		Level:       logrus.InfoLevel,
		Input:       img,
		ParentStack: 1,
	})
}
