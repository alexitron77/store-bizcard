package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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

	err = c.SaveUploadedFile(file, "uploaded/"+file.Filename)

	if err != nil {
		b.config.Log.Errorf(err.Error())
	}

	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	// go utils.Ocr(file, ocr_channel)
	// WriteToWs(<-ocr_channel)
	b.awsRepo.UploadToS3(b.config.S3, file)
}
