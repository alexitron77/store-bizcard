package controllers

import (
	"biz.card/cmd/api/models"
	"biz.card/config"
)

type BizcardController struct {
	config      *config.Logger
	storage     *config.Storage
	bizcardRepo models.BizcardRepo
	awsRepo     models.AwsRepo
}

func NewBizcardController(log *config.Logger, storage *config.Storage, bizcardRepo models.BizcardRepo, awsRepo models.AwsRepo) *BizcardController {
	return &BizcardController{
		log,
		storage,
		bizcardRepo,
		awsRepo,
	}
}
