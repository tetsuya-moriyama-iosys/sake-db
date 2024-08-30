package liquor

import (
	"backend/db"
	"backend/graph/model"
	"backend/util/amazon/s3"
	"backend/util/helper"
	"bytes"
	"context"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Name        string  `form:"title"`
	CategoryID  int     `form:"category"`
	Description *string `form:"description"`
}

func Post(c *gin.Context) {
	var request RequestData
	var imageBase64 *string = nil
	var imageUrl *string = nil

	// 画像以外のフォームデータを構造体にバインド
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// フォームからファイルを取得
	img, _, err := c.Request.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			// ファイルが存在しない場合の処理（画像がアップロードされていない）
			// 例えば、画像が必須ではない場合はこのまま進めることができます
			img = nil
		} else {
			// その他のエラーの場合
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process image file"})
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//画像登録処理
	if img != nil {
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read image file"})
				return
			}
		}(img)

		// 画像データをデコードして、ImageData構造体に格納
		img, format, err := image.Decode(img)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image"})
			return
		}

		imageData := &s3.ImageData{
			Image:  img,
			Format: format,
		}

		// リサイズ実行
		thumbnail := helper.ResizeImage(img, 120, 200)

		// Base64エンコード
		var thumbBuf bytes.Buffer
		err = jpeg.Encode(&thumbBuf, thumbnail, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to encode thumbnail"})
			return
		}

		//string型を*stringに変換する
		encoded := base64.StdEncoding.EncodeToString(thumbBuf.Bytes())
		imageBase64 = &encoded

		imageUrl, err = s3.UploadLiquorImage(imageData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upload image"})
			return
		}
	}

	liquor := &model.Liquor{
		ID:          primitive.NewObjectID().Hex(),
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Description: request.Description,
		ImageURL:    imageUrl,
		ImageBase64: imageBase64,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = db.GetCollection("liquors").InsertOne(ctx, liquor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating liquor"})
		return
	}

	// デバッグ用に受信データをログ出力
	log.Printf("Received liquor data: %+v\n", liquor)

}
