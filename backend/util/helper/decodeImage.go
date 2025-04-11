package helper

import (
	"backend/middlewares/customError"
	"github.com/sirupsen/logrus"
	"image"
	_ "image/png" // PNGデコーダーのインポート
	"mime/multipart"
	"net/http"
)

func DecodeImage(img multipart.File) (image.Image, string, *customError.Error) {
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(img)

	// 画像データをデコードして、ImageData構造体に格納
	result, format, err := image.Decode(img)
	if err != nil {
		return nil, "", errDecodeImage(err, img)
	}

	return result, format, nil
}

func errDecodeImage(err error, img multipart.File) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    "ERR-DecodeImage",
		UserMsg:    "ファイルが不正です",
		Level:      logrus.InfoLevel,
		Input:      img,
	})
}
