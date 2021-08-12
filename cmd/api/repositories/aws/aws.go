package aws

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsRepo struct {
}

func AwsInit(accessKey string, secret string) *s3.S3 {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
		Region:      aws.String("ap-southeast-1"),
	}

	newSession, err := session.NewSession(s3Config)

	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.New(newSession)
	return s3Client
}

func NewAwsRepo() *AwsRepo {
	return &AwsRepo{}
}

func (awsService *AwsRepo) UploadToS3(s3Client *s3.S3, file *multipart.FileHeader) error {
	bucket := aws.String("bizcards")
	key := aws.String(file.Filename)

	fileToUpload, _ := file.Open()

	blob := &s3.PutObjectInput{
		Body:   fileToUpload,
		Bucket: bucket,
		Key:    key,
	}

	_, err := s3Client.PutObject(blob)

	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n", *bucket, *key, err.Error())
		return err
	}

	fmt.Printf("Successfully created bucket %s and uploaded data with key %s\n", *bucket, *key)
	return nil
}
