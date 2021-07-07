package controllers

import (
	"biz.card/models"
	"github.com/gin-gonic/gin"
)

type BizcardController struct {
	bizcardRepo models.BizcardRepo
}

func NewBizcardController(bizcardRepo models.BizcardRepo) *BizcardController {
	return &BizcardController{
		bizcardRepo,
	}
}

func (b *BizcardController) SaveBizCard(c *gin.Context) {
	b.bizcardRepo.Save()
}
