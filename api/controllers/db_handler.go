package controllers

import (
	"biz.card/models"
	"github.com/gin-gonic/gin"
)

type DBHandler struct {
	bizcardRepo models.BizcardRepository
}

func NewBizcardHandler(bizcardRepo models.BizcardRepository) *DBHandler {
	return &DBHandler{
		bizcardRepo,
	}
}

func (r *DBHandler) SaveBizCard(c *gin.Context) {
	r.bizcardRepo.Save()
}
