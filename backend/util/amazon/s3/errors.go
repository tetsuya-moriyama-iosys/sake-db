package s3

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	ReadEnvFailure                = "S3-001-ReadEnvFailure"
	ReadAWSRegionFailure          = "S3-002-ReadAWSRegionFailure"
	ReadAWSImageBucketNameFailure = "S3-003-ReadAWSImageBucketNameFailure"
	CreateAWSImageUploaderFailure = "S3-004-CreateAWSImageUploaderFailure"
	UploadAWSImageFailure         = "S3-005-UploadAWSImageFailure"
	FailureImageEncodeForAWS      = "S3-006-FailureImageEncodeForAWS"
)

func errReadEnvFailure(err error, image *ImageData) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode:  http.StatusInternalServerError,
		ErrCode:     ReadEnvFailure,
		UserMsg:     errorMsg.SERVER,
		Level:       logrus.FatalLevel,
		Input:       *image,
		ParentStack: 1,
	})
}

func errReadAWSRegionFailure() *customError.Error {
	return customError.NewError(errors.New("AWS_REGION environment variable is required"), customError.Params{
		StatusCode:  http.StatusInternalServerError,
		ErrCode:     ReadAWSRegionFailure,
		UserMsg:     errorMsg.SERVER,
		Level:       logrus.FatalLevel,
		ParentStack: 1,
	})
}

func errReadAWSImageBucketNameFailure() *customError.Error {
	return customError.NewError(errors.New("AWS_IMAGE_BUCKET_NAME environment variable is required"), customError.Params{
		StatusCode:  http.StatusInternalServerError,
		ErrCode:     ReadAWSImageBucketNameFailure,
		UserMsg:     errorMsg.SERVER,
		Level:       logrus.FatalLevel,
		ParentStack: 1,
	})
}

func errCreateAWSImageUploaderFailure(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode:  http.StatusInternalServerError,
		ErrCode:     CreateAWSImageUploaderFailure,
		UserMsg:     errorMsg.SERVER,
		Level:       logrus.ErrorLevel,
		ParentStack: 1,
	})
}

func errUploadAWSImageFailure(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode:  http.StatusInternalServerError,
		ErrCode:     UploadAWSImageFailure,
		UserMsg:     errorMsg.SERVER,
		Level:       logrus.ErrorLevel,
		ParentStack: 1,
	})
}

func errFailureImageEncodeForAWS(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode:  http.StatusInternalServerError,
		ErrCode:     FailureImageEncodeForAWS,
		UserMsg:     errorMsg.SERVER,
		Level:       logrus.ErrorLevel,
		ParentStack: 1,
	})
}
