package controllers

import (
	"biz.card/models"
	"github.com/sirupsen/logrus"
)

type BizcardController struct {
	bizcardRepo models.BizcardRepo
	log         *logrus.Logger
}

func NewBizcardController(bizcardRepo models.BizcardRepo, log *logrus.Logger) *BizcardController {
	return &BizcardController{
		bizcardRepo,
		log,
	}
}
