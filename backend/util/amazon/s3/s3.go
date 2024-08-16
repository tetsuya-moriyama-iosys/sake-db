package s3

import (
	"bytes"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Uploader はS3にファイルをアップロードするための構造体
type Uploader struct {
	svc        *s3.S3 //S3のインスタンス
	bucketName string
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
func (u *Uploader) UploadFile(upload *graphql.Upload) (string, error) {
	// ファイルを開く
	file := upload.File
	// fileがio.ReadCloserにキャストできるか確認
	if closer, ok := file.(io.ReadCloser); ok {
		defer func() {
			if err := closer.Close(); err != nil {
				log.Printf("ファイルを閉じるのに失敗しました: %v", err)
			}
		}()
	}

	// ファイル全体をバッファに読み込む
	var buf bytes.Buffer
	fileSize, err := buf.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("画像データの読み込みに失敗しました: %w", err)
	}

	// UUIDを使って一意なファイル名を生成
	ext := filepath.Ext(upload.Filename)
	uniqueFileName := uuid.New().String() + ext

	// S3にファイルをアップロード
	_, err = u.svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(u.bucketName),
		Key:           aws.String(uniqueFileName),
		Body:          bytes.NewReader(buf.Bytes()),
		ContentLength: aws.Int64(fileSize),
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

func UploadLiquorImage(image *graphql.Upload) (*string, error) {
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
