package helper

import (
	"bytes"
	"fmt"
	"image"
	"io"
)

func GetBufImage(image io.ReadSeeker) (*bytes.Buffer, error) {
	// 画像データをバイトスライスとして読み取る
	var buf bytes.Buffer
	_, err := io.Copy(&buf, image)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %v", err)
	}

	return &buf, nil
}

func DecodeImage(buf *bytes.Buffer) (image.Image, error) {
	img, _, err := image.Decode(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}
	return img, nil
}

func GetImageFromRequest(image io.ReadSeeker) (image.Image, error) {
	buf, err := GetBufImage(image)
	if err != nil {
		return nil, err
	}
	img, err := DecodeImage(buf)
	if err != nil {
		return nil, err
	}
	return img, nil
}
