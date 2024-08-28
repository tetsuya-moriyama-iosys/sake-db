package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/graph/model"
	"context"
	"fmt"
)

// CreateLiquor is the resolver for the createLiquor field.
func (r *mutationResolver) CreateLiquor(ctx context.Context, inputs model.CreateLiquorRequest) (*model.Liquor, error) {
	panic(fmt.Errorf("not implemented: Liquors - liquors"))
	//// リクエスト内容をログ出力（デバッグ用）
	//log.Printf("CreateLiquor called with inputs: %+v\n", inputs)
	//
	//var imageUrl *string = nil
	//var imageBase64 *string = nil
	//now := time.Now()
	//
	////画像登録処理
	//if inputs.Image != nil {
	//	log.Println("画像処理開始")
	//	img, err := helper.GetImageFromRequest(inputs.Image.File)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to read image file: %v", err)
	//	}
	//
	//	// リサイズ実行
	//	thumbnail := helper.ResizeImage(img, 120, 200)
	//
	//	// Base64エンコード
	//	var thumbBuf bytes.Buffer
	//	err = jpeg.Encode(&thumbBuf, thumbnail, nil)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to encode thumbnail: %v", err)
	//	}
	//
	//	//string型を*stringに変換する
	//	encoded := base64.StdEncoding.EncodeToString(thumbBuf.Bytes())
	//	imageBase64 = &encoded
	//
	//	imageUrl, err = s3.UploadLiquorImage(inputs.Image)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to upload image: %v", err)
	//	}
	//}
	//
	//liquor := &model.Liquor{
	//	ID:          primitive.NewObjectID().Hex(),
	//	CategoryID:  inputs.CategoryID,
	//	Name:        inputs.Name,
	//	Description: inputs.Description,
	//	ImageURL:    imageUrl,
	//	ImageBase64: imageBase64,
	//	CreatedAt:   now,
	//	UpdatedAt:   now,
	//}
	//
	////DBに挿入
	//_, err := collections.Liquors().InsertOne(ctx, *liquor)
	//if err != nil {
	//	log.Printf("Error inserting liquor: %v", err) // エラーハンドリングの追加
	//	return nil, err
	//}
	//
	//log.Printf("Liquor successfully created: %+v\n", liquor)
	//return liquor, nil
}

// Liquors is the resolver for the liquors field.
func (r *queryResolver) Liquors(ctx context.Context) ([]*model.Liquor, error) {
	panic(fmt.Errorf("not implemented: Liquors - liquors"))
}

// Liquor is the resolver for the liquor field.
func (r *queryResolver) Liquor(ctx context.Context, id int) (*model.Liquor, error) {
	panic(fmt.Errorf("not implemented: Liquor - liquor"))
}
