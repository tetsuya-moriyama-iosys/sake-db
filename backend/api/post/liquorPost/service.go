package liquorPost

import (
	"backend/const/errorMsg"
	"backend/db"
	"backend/db/repository/liquorRepository"
	"backend/util/amazon/s3"
	"backend/util/helper"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Id          string  `form:"id"`
	Name        string  `form:"name"`
	CategoryID  int     `form:"category"`
	Description *string `form:"description"`
	VersionNo   *int    `form:"version_no"`
}

// Base64にリサイズする際の横幅
var maxWidth uint = 200

func (h *Handler) Post(c *gin.Context) (*string, error) {
	var request RequestData
	var imageBase64 *string
	var imageUrl *string

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 画像以外のフォームデータを構造体にバインド
	if err := c.ShouldBind(&request); err != nil {
		return nil, errors.New("invalid form data")
	}

	//logsに代入する現在のドキュメントを取得する
	old, err := h.LiquorsRepo.GetLiquorById(ctx, request.Id)

	//IDが空でない(=編集)の場合、バージョンnoのチェックを行う
	if !helper.IsEmpty(&request.Id) {
		if old.VersionNo != *request.VersionNo {
			return nil, errors.New(errorMsg.VERSION)
		}
	}

	// フォームからファイルを取得
	rawImg, _, err := c.Request.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			// 画像が存在しない場合
			rawImg = nil
		} else {
			// その他のエラーの場合
			return nil, errors.New("failed to process image file")
		}
	}

	//画像登録処理
	if rawImg != nil {
		// 画像データをデコード
		img, format, err := helper.DecodeImage(rawImg)
		if err != nil {
			return nil, errors.New("failed to decode image")
		}

		//base64エンコードしたデータを取得
		maxHeight := maxWidth / 9 * 16
		imageBase64, err = helper.ImageToBase64(img, &helper.Base64Option{
			MaxWidth:  &maxWidth,
			MaxHeight: &maxHeight,
		})
		if err != nil {
			return nil, errors.New("failed to convert image to base64")
		}

		//S3にアップロードし、URLを取得する
		imageUrl, err = s3.UploadLiquorImage(&s3.ImageData{
			Image:  img,
			Format: format,
		})
		if err != nil {
			return nil, errors.New("failed to upload image")
		}
	}

	//カテゴリ名を取得する
	category, err := h.CategoryRepo.GetCategoryByID(ctx, request.CategoryID)
	if err != nil {
		return nil, err
	}

	//新バージョンNoを作成する
	var newVersionNo int
	var newCreatedAt time.Time
	var id primitive.ObjectID
	if request.VersionNo != nil {
		//更新の場合
		id, _ = primitive.ObjectIDFromHex(request.Id) //エラーが出ることはここでは考慮しない
		newVersionNo = *request.VersionNo + 1
		newCreatedAt = old.CreatedAt
	} else {
		//初回作成の場合
		id = primitive.NewObjectID()
		newVersionNo = 1 // 初回作成の場合、VersionNoを1に設定
		newCreatedAt = time.Now()
	}

	//画像は毎回送信しないため、フォームが空であれば前回の値をそのまま代入
	var newBase64 *string
	var newImageURL *string
	if rawImg != nil {
		newBase64 = imageBase64
		newImageURL = imageUrl
	} else {
		newBase64 = old.ImageBase64
		newImageURL = old.ImageURL
	}

	//挿入するドキュメントを作成
	record := &liquorRepository.Model{
		ID:           id,
		CategoryID:   request.CategoryID,
		CategoryName: category.Name,
		Name:         request.Name,
		Description:  request.Description,
		ImageURL:     newImageURL,
		ImageBase64:  newBase64,
		CreatedAt:    newCreatedAt,
		UpdatedAt:    time.Now(),
		VersionNo:    newVersionNo,
	}

	//トランザクション
	newId, err := db.WithTransaction(ctx, h.DB.Client(), func(sc mongo.SessionContext) (*string, error) {
		// トランザクション内での操作1
		if old == nil {
			//新規追加
			newObjId, err := h.LiquorsRepo.InsertOne(ctx, record)
			if err != nil {
				return nil, err
			}
			newObjIdStr := newObjId.Hex()
			return &newObjIdStr, nil
		}
		//更新
		newObjId, err := h.LiquorsRepo.UpdateOne(ctx, record)
		if err != nil {
			return nil, err
		}

		//logsに追加
		if !helper.IsEmpty(&request.Id) {
			err = h.LiquorsRepo.InsertOneToLog(ctx, old)
			if err != nil {
				return nil, err
			}
		}
		newObjIdStr := newObjId.Hex()
		return &newObjIdStr, nil
	})
	if err != nil {
		return nil, err
	}
	return newId, nil
}
