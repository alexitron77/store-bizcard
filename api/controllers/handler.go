package controllers

import (
	"biz.card/models"
)

type BizcardController struct {
	bizcardRepo models.BizcardRepo
}

func NewBizcardController(bizcardRepo models.BizcardRepo) *BizcardController {
	return &BizcardController{
		bizcardRepo,
	}
}
