package controllers

import (
	"net/http"

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
	var card models.Bizcard
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err := b.bizcardRepo.Save(&card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Card successfully created"})
}
