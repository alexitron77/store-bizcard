package config

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logger struct {
	Log *logrus.Logger
}

type Storage struct {
	DB *mongo.Client
	S3 *s3.S3
}

func NewLogger(log *logrus.Logger) *Logger {
	return &Logger{
		Log: log,
	}
}

func NewStorage(db *mongo.Client, s3 *s3.S3) *Storage {
	return &Storage{
		DB: db,
		S3: s3,
	}
}
