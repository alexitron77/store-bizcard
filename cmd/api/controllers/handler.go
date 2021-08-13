package controllers

import (
	"biz.card/cmd/api/models"
	"biz.card/config"
)

type BizcardController struct {
	config      *config.Config
	storage     *config.Storage
	bizcardRepo models.BizcardRepo
	awsRepo     models.AwsRepo
}

func NewBizcardController(config *config.Config, storage *config.Storage, bizcardRepo models.BizcardRepo, awsRepo models.AwsRepo) *BizcardController {
	return &BizcardController{
		config,
		storage,
		bizcardRepo,
		awsRepo,
	}
}
