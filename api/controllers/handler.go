package controllers

import (
	"biz.card/config"
	"biz.card/models"
)

type BizcardController struct {
	config      *config.Config
	bizcardRepo models.BizcardRepo
	awsRepo     models.AwsRepo
}

func NewBizcardController(config *config.Config, bizcardRepo models.BizcardRepo, awsRepo models.AwsRepo) *BizcardController {
	return &BizcardController{
		config,
		bizcardRepo,
		awsRepo,
	}
}
