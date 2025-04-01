package categoryPost

import (
	"backend/db"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/userRepository"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"backend/service/categoryService"
	"backend/util/amazon/s3"
	"backend/util/helper"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func (h *Handler) Post(c *gin.Context, ur *userRepository.UsersRepository) (*int, *customError.Error) {
	ctx := c.Request.Context()

	var request RequestData
	var imageBase64 *string
	var imageUrl *string
	var old *categoriesRepository.Model

	uId, uName, err := auth.GetIdAndNameNullable(c, ur)
	if err != nil {
		return nil, err
	}

	// 画像以外のフォームデータを構造体にバインド
	if err := c.ShouldBind(&request); err != nil {
		return nil, errInvalidInput(c, err)
	}

	if request.Id != nil {
		//更新時のみ行う処理
		hasIdInTrail, err := categoryService.HasIdInTrail(ctx, &h.CategoryRepo, *request.Id, request.Parent)
		if err != nil {
			return nil, err
		}
		if hasIdInTrail == true {
			return nil, errInvalidParent(request)
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
			return nil, errInvalidVersion(request)
		}
	}

	// フォームからファイルを取得
	rawImg, _, fileErr := c.Request.FormFile("image")
	if fileErr != nil {
		if errors.Is(fileErr, http.ErrMissingFile) {
			// 画像が存在しない場合
			rawImg = nil
		} else {
			// その他のエラーの場合
			return nil, errInvalidFile(fileErr, request)
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
		UserId:      uId,
		UserName:    uName,
		UpdatedAt:   time.Now(),
		VersionNo:   &newVersionNo,
	}

	//トランザクション
	// TODO:トランザクションはレプリカセットを使わないと効かないが、ローカルだと構築するのが厳しかったので MongoDB Atlas Databaseの利用を前提に考えることにした
	// NOTE: customError型がerrorに型推論できないようなので、一旦代入して手動でキャストして対応する
	_, iErr := db.WithTransaction(ctx, h.DB.Client(), func(sc mongo.SessionContext) (struct{}, error) {
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
	if iErr != nil {
		errors.As(iErr, &err)
	}
	if err != nil {
		return nil, err
	}
	return &record.ID, nil
}
