package liquorPost

import (
	"backend/graph/graphModel/liquor"
	"backend/util/amazon/s3"
	"backend/util/helper"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Name        string  `form:"title"`
	CategoryID  int     `form:"category"`
	Description *string `form:"description"`
}

// Base64にリサイズする際の横幅
var maxWidth uint = 200

func (h *Handler) Post(c *gin.Context) {
	var request RequestData
	var imageBase64 *string
	var imageUrl *string

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 画像以外のフォームデータを構造体にバインド
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// フォームからファイルを取得
	rawImg, _, err := c.Request.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			// 画像が存在しない場合
			rawImg = nil
		} else {
			// その他のエラーの場合
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process image file"})
			return
		}
	}

	//画像登録処理
	if rawImg != nil {
		// 画像データをデコード
		img, format, err := helper.DecodeImage(rawImg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image"})
			return
		}

		//base64エンコードしたデータを取得
		maxHeight := maxWidth / 9 * 16
		imageBase64, err = helper.ImageToBase64(img, &helper.Base64Option{
			MaxWidth:  &maxWidth,
			MaxHeight: &maxHeight,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert image to base64"})
			return
		}

		//S3にアップロードし、URLを取得する
		imageUrl, err = s3.UploadLiquorImage(&s3.ImageData{
			Image:  img,
			Format: format,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upload image"})
			return
		}
	}

	//カテゴリ名を取得する
	name, err := h.CategoryRepo.GetCategoryNameByID(ctx, request.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := &liquorModel.Schema{
		ID:           primitive.NewObjectID(),
		CategoryID:   request.CategoryID,
		CategoryName: name,
		Name:         request.Name,
		Description:  request.Description,
		ImageURL:     imageUrl,
		ImageBase64:  imageBase64,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err = h.LiquorsRepo.InsertOne(ctx, record)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert failed"})
		return
	}
}
