package helper

import (
	"github.com/nfnt/resize"
	"image"
)

func ResizeImage(img image.Image, maxWidth uint, maxHeight uint) image.Image {
	// 画像の幅と高さを取得
	origWidth := img.Bounds().Dx()
	origHeight := img.Bounds().Dy()

	// アスペクト比を維持しつつ、指定された領域に収めるための計算
	ratio := float64(origWidth) / float64(origHeight)

	var newWidth, newHeight uint

	if uint(origWidth) > maxWidth || uint(origHeight) > maxHeight {
		// 画像が指定された領域を超える場合
		if float64(maxWidth)/ratio <= float64(maxHeight) {
			newWidth = maxWidth
			newHeight = uint(float64(maxWidth) / ratio)
		} else {
			newHeight = maxHeight
			newWidth = uint(float64(maxHeight) * ratio)
		}
	} else {
		// 画像が指定された領域内に収まる場合、そのままのサイズ
		newWidth = uint(origWidth)
		newHeight = uint(origHeight)
	}

	// リサイズ実行
	return resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
}
