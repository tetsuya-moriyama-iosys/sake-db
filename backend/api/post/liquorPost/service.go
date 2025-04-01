package liquorPost

import (
	"backend/db"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"backend/middlewares/auth"
	"backend/middlewares/customError"
	"backend/util/amazon/s3"
	"backend/util/helper"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"net/http"
	"time"
)

func (h *Handler) Post(c *gin.Context, ur *userRepository.UsersRepository) (*string, *customError.Error) {
	ctx := c.Request.Context()

	var request RequestData
	var imageBase64 *string
	var imageUrl *string
	var old *liquorRepository.Model

	uId, uName, err := auth.GetIdAndNameNullable(c, ur)
	if err != nil {
		return nil, err
	}

	// 画像以外のフォームデータを構造体にバインド
	if err := c.ShouldBind(&request); err != nil {
		return nil, errInvalidInput(c, err)
	}

	var id *primitive.ObjectID
	if request.Id != nil {
		tempId, err := primitive.ObjectIDFromHex(*request.Id)
		if err != nil {
			return nil, errParseTempID(err)
		}
		id = &tempId
	}

	//名前の重複チェックを行う
	l, err := h.LiquorsRepo.GetLiquorByName(ctx, request.Name, id)
	if err != nil && !errors.Is(err.RawErr, mongo.ErrNoDocuments) {
		//見つからないエラーは正常系だが、それ以外のエラーの場合
		return nil, err
	}
	if l != nil {
		return nil, errDuplicateName()
	}

	if request.Id != nil {
		//更新時のみ行う処理
		lId, HexErr := primitive.ObjectIDFromHex(*request.Id)
		if err != nil {
			return nil, errParseID(HexErr)
		}
		//logsに代入する現在のドキュメントを取得する
		old, err = h.LiquorsRepo.GetLiquorById(ctx, lId)
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
			return nil, errInvalidVersion()
		}
	}

	// フォームからファイルを取得
	rawImg, _, fErr := c.Request.FormFile("image")
	if fErr != nil {
		if errors.Is(fErr, http.ErrMissingFile) {
			// 画像が存在しない場合
			rawImg = nil
		} else {
			// その他のエラーの場合
			return nil, errInvalidFile(fErr, rawImg)
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
		imgOld, err := h.LiquorsRepo.GetLogsByVersionNo(ctx, *request.Id, *request.SelectedVersionNo)
		if err != nil {
			return nil, err
		}
		old.ImageBase64 = imgOld.ImageBase64
		old.ImageURL = imgOld.ImageURL
	}

	//カテゴリ名を取得する
	category, err := h.CategoryRepo.GetCategoryByID(ctx, request.CategoryID)
	if err != nil {
		return nil, err
	}

	//新バージョンNoを作成する
	var newVersionNo int
	if id != nil {
		//更新の場合
		if request.VersionNo == nil {
			//初期アセットの場合(version_noを入れていない)
			newVersionNo = 1
		} else {
			newVersionNo = *request.VersionNo + 1
		}
	} else {
		//初回作成の場合
		tempId := primitive.NewObjectID()
		id = &tempId
		newVersionNo = 1 // 初回作成の場合、VersionNoを1に設定
	}

	//画像は毎回送信しないため、フォームが空であれば前回の値をそのまま代入
	var newBase64 *string
	var newImageURL *string
	if rawImg != nil {
		newBase64 = imageBase64
		newImageURL = imageUrl
	} else {
		if old != nil {
			newBase64 = old.ImageBase64
			newImageURL = old.ImageURL
		}
	}

	//挿入するドキュメントを作成
	record := &liquorRepository.Model{
		ID:           *id,
		CategoryID:   request.CategoryID,
		CategoryName: category.Name,
		Name:         request.Name,
		Description:  &request.Description,
		Youtube:      &request.Youtube,
		ImageURL:     newImageURL,
		ImageBase64:  newBase64,
		UpdatedAt:    time.Now(),
		RandomKey:    rand.New(rand.NewSource(time.Now().UnixNano())).Float64(), //毎回更新する
		VersionNo:    &newVersionNo,
		UserId:       uId,
		UserName:     uName,
	}

	//トランザクション
	newId, iErr := db.WithTransaction(ctx, h.DB.Client(), func(sc mongo.SessionContext) (*string, error) {
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

	if iErr != nil {
		errors.As(iErr, &err)
	}
	if err != nil {
		return nil, err
	}
	return newId, nil
}
