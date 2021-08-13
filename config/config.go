package config

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	DB  *mongo.Client
	Log *logrus.Logger
	S3  *s3.S3
}

func NewConfig(db *mongo.Client, log *logrus.Logger, s3 *s3.S3) *Config {
	return &Config{
		DB:  db,
		Log: log,
		S3:  s3,
	}
}
