package amazon

import (
	"bytes"
)

// Amazon S3に画像をアップロードする関数
func UploadToS3(filename string, fileData bytes.Buffer) (string, error) {
	// S3クライアントの初期化（設定は省略）
	//s3Client := s3.NewFromConfig(aws.Config{ /* S3の設定をここに記載 */ })
	//
	//uploader := manager.NewUploader(s3Client)
	//
	//// S3にアップロード
	//result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
	//	Bucket: aws.String("your-bucket-name"),
	//	Key:    aws.String(filename),
	//	Body:   bytes.NewReader(fileData.Bytes()),
	//})
	//if err != nil {
	//	return "", fmt.Errorf("failed to upload to S3: %v", err)
	//}
	//
	//return result.Location, nil
}
