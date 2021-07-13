package controllers

import (
	"net/http"

	"biz.card/models"
	"github.com/gin-gonic/gin"
)

// SaveCard godoc
// @Summary Save card
// @Description This endpoint save the input into the database
// @ID save-card-to-database
// @Accept json
// @Produce json
// @Param card body models.Bizcard true "Create bizcard"
// @Success 201 {object} models.HTTPCreated
// @Failure 400 {object} models.HTTPClientError
// @Failure 500 {object} models.HTTPBackendError
// @Router /create-card [post]
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
