package config

import (
	"context"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	DB  *mongo.Client
	Ctx context.Context
	Log *logrus.Logger
	S3  *s3.S3
}

func NewConfig(db *mongo.Client, ctx context.Context, log *logrus.Logger, s3 *s3.S3) *Config {
	return &Config{
		DB:  db,
		Ctx: ctx,
		Log: log,
		S3:  s3,
	}
}
