package controllers

import (
	"fmt"
	"net/http"
	"os"

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
		b.config.Log.Errorf(err.Error())
		return
	}

	err := b.bizcardRepo.Create(&card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		b.config.Log.Errorf(err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Card successfully created"})
}

// Upload godoc
// @Summary Upload file
// @Description This endpoint upload an image file into the file system of the server
// @ID upload-file-to-server
// @Accept json
// @Produce json
// @Param myFile formData file true "Body with image file"
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPClientError
// @Failure 500 {object} models.HTTPBackendError
// @Router /upload-card [post]
func (b *BizcardController) Upload(c *gin.Context) {
	// ocr_channel := make(chan string)

	file, err := c.FormFile("myFile")

	if err != nil {
		b.config.Log.Errorf(err.Error())
		os.Exit(1)
	}

	// go utils.Ocr(file, ocr_channel)
	// WriteToWs(<-ocr_channel)

	err = b.awsRepo.UploadToS3(b.config.S3, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Upload failed! Err:%s", err))
	}

	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (b *BizcardController) ReadBizCard(c *gin.Context) {
	id := c.Param("name")

	res, err := b.bizcardRepo.Read(id)

	if err != nil {
		b.config.Log.Errorf(err.Error())
	}

	c.JSON(http.StatusOK, res)
}
