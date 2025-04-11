package s3

import (
	"backend/middlewares/customError"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"image"
	"image/jpeg"
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

func NewS3Client() (*s3.S3, error) {
	awsRegion := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}

// NewS3Uploader はS3Uploaderを初期化するためのファクトリ関数
func NewS3Uploader(region, bucketName string) (*Uploader, *customError.Error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, errCreateAWSImageUploaderFailure(err)
	}

	svc := s3.New(sess)
	return &Uploader{
		svc:        svc,
		bucketName: bucketName,
	}, nil
}

// UploadFile UploadFileFromForm はHTTPリクエストのフォームから取得したファイルをS3にアップロードし、アップロードされたファイルのURLを返します
func (u *Uploader) UploadFile(image *ImageData) (string, *customError.Error) {

	// 画像データをJPEGとしてエンコードするためにバッファに書き込む
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, image.Image, nil)
	if err != nil {
		return "", errFailureImageEncodeForAWS(err)
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
		return "", errUploadAWSImageFailure(err)
	}

	// アップロードされたファイルのURLを返す
	return u.getFileURL(uniqueFileName), nil
}

// getFileURL は指定されたキーに対応するS3オブジェクトのURLを返します
func (u *Uploader) getFileURL(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.bucketName, *u.svc.Config.Region, key)
}

func UploadLiquorImage(image *ImageData) (*string, *customError.Error) {
	// .envファイルを読み込みます
	envErr := godotenv.Load()
	if envErr != nil {
		return nil, errReadEnvFailure(envErr, image)
	}

	// 環境変数から必要な値を取得
	region := os.Getenv("AWS_REGION")
	if region == "" {
		return nil, errReadAWSRegionFailure()
	}
	bucket := os.Getenv("AWS_IMAGE_BUCKET_NAME")
	if bucket == "" {
		return nil, errReadAWSImageBucketNameFailure()
	}

	// S3アップロード処理
	uploader, err := NewS3Uploader(region, bucket)
	if err != nil {
		return nil, err
	}

	s3Url, err := uploader.UploadFile(image)
	if err != nil {
		return nil, err
	}
	return &s3Url, nil
}
