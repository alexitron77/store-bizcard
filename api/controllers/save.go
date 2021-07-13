package controllers

import (
	"net/http"

	"biz.card/models"
	"github.com/gin-gonic/gin"
)

func (b *BizcardController) SaveBizCard(c *gin.Context) {
	var card models.Bizcard
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		b.log.Errorf(err.Error())
		return
	}

	err := b.bizcardRepo.Save(&card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		b.log.Errorf(err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Card successfully created"})
}
