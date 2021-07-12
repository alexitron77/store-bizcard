package controllers

import (
	"fmt"
	"log"
	"net/http"

	"biz.card/models"
	"biz.card/utils"
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

func (b *BizcardController) Upload(c *gin.Context) {
	ocr_channel := make(chan string)

	file, err := c.FormFile("myFile")

	if err != nil {
		log.Fatal(err)
	}

	err = c.SaveUploadedFile(file, "uploaded/"+file.Filename)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Print(err)
	}

	go utils.Ocr(file, ocr_channel)
	fmt.Print(<-ocr_channel)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
