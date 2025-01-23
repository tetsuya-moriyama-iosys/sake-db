package helper

import (
	"fmt"
	"image"
	"net/http"
)

func FetchImageFromURL(url string) (image.Image, error) {
	// URLから画像データを取得
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// HTTPステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image, status code: %d", resp.StatusCode)
	}

	// 画像データをデコード
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
