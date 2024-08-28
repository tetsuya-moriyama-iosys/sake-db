package s3

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

// Uploader はS3にファイルをアップロードするための構造体
type Uploader struct {
	svc        *s3.S3 //S3のインスタンス
	bucketName string
}

type ImageData struct {
	Image  image.Image // デコードされた画像データ
	Format string      // 画像のフォーマット（例: "jpeg", "png" など）
}

// NewS3Uploader はS3Uploaderを初期化するためのファクトリ関数
func NewS3Uploader(region, bucketName string) (*Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("セッションの作成に失敗しました: %w", err)
	}

	svc := s3.New(sess)
	return &Uploader{
		svc:        svc,
		bucketName: bucketName,
	}, nil
}

// UploadFile UploadFileFromForm はHTTPリクエストのフォームから取得したファイルをS3にアップロードし、アップロードされたファイルのURLを返します
func (u *Uploader) UploadFile(image *ImageData) (string, error) {

	// 画像データをJPEGとしてエンコードするためにバッファに書き込む
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, image.Image, nil)
	if err != nil {
		return "", fmt.Errorf("画像データのエンコードに失敗しました: %w", err)
	}

	// UUIDを使って一意なファイル名を生成
	uniqueFileName := uuid.New().String() + image.Format

	// S3にファイルをアップロード
	_, err = u.svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(u.bucketName),
		Key:           aws.String(uniqueFileName),
		Body:          bytes.NewReader(buf.Bytes()),
		ContentLength: aws.Int64(int64(buf.Len())),
		ContentType:   aws.String(http.DetectContentType(buf.Bytes())),
	})
	if err != nil {
		return "", fmt.Errorf("画像のアップロードに失敗しました: %w", err)
	}

	// アップロードされたファイルのURLを返す
	return u.getFileURL(uniqueFileName), nil
}

// getFileURL は指定されたキーに対応するS3オブジェクトのURLを返します
func (u *Uploader) getFileURL(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.bucketName, *u.svc.Config.Region, key)
}

func UploadLiquorImage(image *ImageData) (*string, error) {
	// .envファイルを読み込みます
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: loading .env file")
	}

	// 環境変数から必要な値を取得
	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Fatal("Error: AWS_REGION environment variable is required")
	}
	bucket := os.Getenv("AWS_IMAGE_BUCKET_NAME")
	if bucket == "" {
		log.Fatal("Error: AWS_IMAGE_BUCKET_NAME environment variable is required")
	}

	// S3アップロード処理
	uploader, err := NewS3Uploader(region, bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 uploader: %v", err)
	}

	s3Url, err := uploader.UploadFile(image)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	return &s3Url, nil
}
