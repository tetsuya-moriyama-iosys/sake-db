package categoryPost

import (
	"backend/const/errorMsg"
	"backend/db"
	"backend/db/repository/categoriesRepository"
	"backend/service/categoryService"
	"backend/util/amazon/s3"
	"backend/util/helper"
	"backend/util/validator"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func (h *Handler) Post(c *gin.Context) (*int, error) {
	var request RequestData
	var imageBase64 *string
	var imageUrl *string
	var old *categoriesRepository.Model

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 画像以外のフォームデータを構造体にバインド
	if err := c.ShouldBind(&request); err != nil {
		return nil, errors.New("invalid form data")
	}

	err := validator.Validate(request)
	if err != nil {
		return nil, err
	}

	if request.Id != nil {
		//更新時のみ行う処理
		hasIdInTrail, err := categoryService.HasIdInTrail(ctx, &h.CategoryRepo, *request.Id, request.Parent)
		if err != nil {
			return nil, err
		}
		if hasIdInTrail == true {
			return nil, errors.New("自身または子カテゴリを親とすることはできません")
		}

		//logsに代入する現在のドキュメントを取得する
		old, err = h.CategoryRepo.GetCategoryByID(ctx, *request.Id)
		if err != nil {
			return nil, err
		}
		//nil参照エラー回避が面倒なので、nilは0扱いとする(versionNoがスキーマ上後付なので、nilの可能性がある)
		if old.VersionNo == nil {
			zero := 0
			old.VersionNo = &zero
		}
		if request.VersionNo == nil {
			verZero := 0
			request.VersionNo = &verZero
		}
		//旧バージョンno(今あるDBのバージョンno)が空でない場合のみチェックする
		if *old.VersionNo != *request.VersionNo {
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
			return nil, err
		}
	}

	//画像登録処理
	if rawImg != nil {
		// 画像データをデコード
		img, format, err := helper.DecodeImage(rawImg)
		if err != nil {
			return nil, err
		}

		//base64エンコードしたデータを取得
		maxHeight := maxWidth / 9 * 16
		imageBase64, err = helper.ImageToBase64(img, &helper.Base64Option{
			MaxWidth:  &maxWidth,
			MaxHeight: &maxHeight,
		})
		if err != nil {
			return nil, err
		}

		//S3にアップロードし、URLを取得する
		imageUrl, err = s3.UploadLiquorImage(&s3.ImageData{
			Image:  img,
			Format: format,
		})
		if err != nil {
			return nil, err
		}
	} else if request.SelectedVersionNo != nil {
		//画像が存在しないが、選択されたロールバック先がある、つまり画像のロールバックが考えうる
		imgOld, err := h.CategoryRepo.GetLogsByVersionNo(ctx, *request.Id, *request.SelectedVersionNo)
		if err != nil {
			return nil, err
		}
		old.ImageBase64 = imgOld.ImageBase64
		old.ImageURL = imgOld.ImageURL
	}

	//ここから新規・更新で処理を共通にする
	//新バージョンNoを作成する
	var newVersionNo int
	var id int
	if request.Id != nil {
		//更新の場合
		id = *request.Id
		if request.VersionNo == nil {
			//初期アセットの場合(version_noを入れていない)
			newVersionNo = 1
		} else {
			newVersionNo = *request.VersionNo + 1
		}
	} else {
		//初回作成の場合
		maxId, err := h.CategoryRepo.GetMaxID(ctx)
		if err != nil {
			return nil, err
		}
		id = maxId + 1
		newVersionNo = 1 // 初回作成の場合、VersionNoを1に設定
	}

	//画像は毎回送信しないため、フォームが空であれば前回の値をそのまま代入
	var newBase64 *string
	var newImageURL *string
	if rawImg != nil {
		newBase64 = imageBase64
		newImageURL = imageUrl
	} else {
		//画像が更新されなかった場合、旧データがある場合はそこからコピーする
		if old != nil {
			newBase64 = old.ImageBase64
			newImageURL = old.ImageURL
		}
	}

	//挿入するドキュメントを作成
	record := &categoriesRepository.Model{
		ID:          id,
		Parent:      &request.Parent,
		Name:        request.Name,
		Description: request.Description,
		ImageURL:    newImageURL,
		ImageBase64: newBase64,
		UpdatedAt:   time.Now(),
		VersionNo:   &newVersionNo,
	}

	//トランザクション
	//TODO:トランザクションはレプリカセットを使わないと効かないが、ローカルだと構築するのが厳しかったので MongoDB Atlas Databaseの利用を前提に考えることにした
	_, err = db.WithTransaction(ctx, h.DB.Client(), func(sc mongo.SessionContext) (struct{}, error) {
		zero := struct{}{}

		//logsに追加(ログの更新がコケて新規/更新だけが実行される、というパターンの方が最悪なので、ログを優先的に更新する(本来はトランザクションですが...))
		if !helper.IsEmpty(request.Id) {
			err = h.CategoryRepo.InsertOneToLog(ctx, old)
			if err != nil {
				return zero, err
			}
		}

		if old == nil {
			//新規追加
			err := h.CategoryRepo.InsertOne(ctx, record)
			if err != nil {
				return zero, err
			}
			return zero, nil
		}
		//更新
		err := h.CategoryRepo.UpdateOne(ctx, record)
		if err != nil {
			return zero, err
		}

		return zero, nil
	})
	if err != nil {
		return nil, err
	}
	return &record.ID, nil
}
