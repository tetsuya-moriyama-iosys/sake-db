package helper

import (
	"image"
	"mime/multipart"
)

func DecodeImage(img multipart.File) (image.Image, string, error) {
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(img)

	// 画像データをデコードして、ImageData構造体に格納
	result, format, err := image.Decode(img)
	if err != nil {
		return nil, "", err
	}

	return result, format, err
}
