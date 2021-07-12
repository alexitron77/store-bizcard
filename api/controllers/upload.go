package controllers

import (
	"fmt"
	"log"
	"net/http"

	"biz.card/utils"
	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	go utils.Ocr(file, ocr_channel)
	WriteToWs(<-ocr_channel)
}
