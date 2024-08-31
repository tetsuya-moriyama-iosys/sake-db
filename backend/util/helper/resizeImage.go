package helper

import (
	"github.com/nfnt/resize"
	"image"
)

func ResizeImage(img image.Image, maxWidth *uint, maxHeight *uint) image.Image {
	// 画像の幅と高さを取得
	origWidth := img.Bounds().Dx()
	origHeight := img.Bounds().Dy()

	// デフォルト値の設定
	var newWidth, newHeight uint
	if maxWidth == nil {
		newWidth = uint(origWidth)
	} else {
		newWidth = *maxWidth
	}

	if maxHeight == nil {
		newHeight = uint(origHeight)
	} else {
		newHeight = *maxHeight
	}

	// アスペクト比を維持しつつ、指定された領域に収めるための計算
	ratio := float64(origWidth) / float64(origHeight)

	if newWidth < uint(origWidth) || newHeight < uint(origHeight) {
		// 画像が指定された領域を超える場合
		if float64(newWidth)/ratio <= float64(newHeight) {
			newHeight = uint(float64(newWidth) / ratio)
		} else {
			newWidth = uint(float64(newHeight) * ratio)
		}
	}

	// リサイズ実行
	return resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
}
