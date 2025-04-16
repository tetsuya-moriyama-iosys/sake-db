package helper

import (
	"backend/middlewares/customError"
	"errors"
	"github.com/sirupsen/logrus"
	"image"
	"io"
	"net/http"
)

const (
	HttpGet    = "Helper-FetchImageFromURL-httpGet"
	StatusCode = "Helper-FetchImageFromURL-StatusCode"
	Decode     = "Helper-FetchImageFromURL-StatusCode"
)

func FetchImageFromURL(url string) (image.Image, *customError.Error) {
	// URLから画像データを取得
	resp, err := http.Get(url)
	if err != nil {
		return nil, errHttpGet(err, url)
	}
	defer resp.Body.Close()

	// HTTPステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		return nil, errStatusCode(resp.StatusCode)
	}

	// 画像データをデコード
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, errDecode(err, resp.Body)
	}

	return img, nil
}

func errHttpGet(err error, url string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    HttpGet,
		UserMsg:    "画像取得に失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      url,
	})
}
func errStatusCode(code int) *customError.Error {
	return customError.NewError(errors.New("failed to fetch image"), customError.Params{
		StatusCode: code,
		ErrCode:    StatusCode,
		UserMsg:    "画像取得に失敗しました",
		Level:      logrus.InfoLevel,
	})
}
func errDecode(err error, body io.ReadCloser) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    Decode,
		UserMsg:    "画像のデコードに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      body,
	})
}
