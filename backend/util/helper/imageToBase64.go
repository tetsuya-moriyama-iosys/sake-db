package helper

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
)

type Base64Option struct {
	MaxWidth  *uint // リサイズ時の最大幅
	MaxHeight *uint // リサイズ時の高さ
}

func ImageToBase64(img image.Image, option *Base64Option) (*string, error) {
	// リサイズ実行
	thumbnail := ResizeImage(img, option.MaxWidth, option.MaxHeight)

	// Base64エンコード
	var thumbBuf bytes.Buffer
	err := jpeg.Encode(&thumbBuf, thumbnail, nil)
	if err != nil {
		return nil, err
	}

	//string型を*stringに変換する
	encoded := base64.StdEncoding.EncodeToString(thumbBuf.Bytes())
	imageBase64 := &encoded

	return imageBase64, nil
}
