package models

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsRepo interface {
	UploadToS3(s3Client *s3.S3, file *multipart.FileHeader) error
}
