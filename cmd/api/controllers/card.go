package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"biz.card/cmd/api/models"
	"biz.card/cmd/kafka"
	"github.com/gin-gonic/gin"
)

// SaveCard godoc
// @Summary Save card
// @Description This endpoint save the input into the database
// @ID save-card-to-database
// @Accept multipart/form-data
// @Produce json
// @Param card body models.Bizcard true "Create bizcard"
// @Success 201 {object} models.HTTPCreated
// @Failure 400 {object} models.HTTPClientError
// @Failure 500 {object} models.HTTPBackendError
// @Router /create-card [post]
func (b *BizcardController) SaveBizCard(c *gin.Context) {

	var card models.Bizcard
	cCard := c.PostForm("card")

	err := json.Unmarshal([]byte(cCard), &card)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		b.config.Log.Errorf(err.Error())
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Millisecond)

	id, err := b.bizcardRepo.Create(ctx, &card)

	if err != nil {
		fmt.Println("Request failed:", err)
		c.Abort()
		return
	}

	c.Set("cardId", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		b.config.Log.Errorf(err.Error())
		return
	}

	c.Next()

	c.JSON(http.StatusCreated, gin.H{"status": "Card successfully created"})

	kafka.Producer([]string{b.config.Config.Kafka.Url}, b.config.Config.Kafka.Topic, fmt.Sprintf("{'status': 'success', 'cardId': %#v", id))
	go kafka.Consumer([]string{b.config.Config.Kafka.Url}, b.config.Config.Kafka.Topic)
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

	file, err := c.FormFile("file")

	if err != nil {
		b.config.Log.Errorf(err.Error())
		os.Exit(1)
	}

	// go utils.Ocr(file, ocr_channel)
	// WriteToWs(<-ocr_channel)

	err = b.awsRepo.UploadToS3(b.storage.S3, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Upload failed! Err:%s", err))
		return
	}
	c.Next()
}

func (b *BizcardController) UpdateCardURL(c *gin.Context) {
	file, _ := c.FormFile("file")

	amazonUrl := fmt.Sprintf("https://bizcards.s3.ap-southeast-1.amazonaws.com/%s", file.Filename)

	id := c.GetString("cardId")
	val := amazonUrl
	ctx, _ := context.WithTimeout(c.Request.Context(), 2000*time.Millisecond)

	b.bizcardRepo.Update(ctx, id, val)
}

// @Summary Read card from DB
// @Description This endpoint retrieve a card from the database
// @ID read-card-from-db
// @Accept  json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPClientError
// @Failure 500 {object} models.HTTPBackendError.
// @Router /get-card/{id} [get]
func (b *BizcardController) ReadBizCard(c *gin.Context) {
	id := c.Param("id")

	res, err := b.bizcardRepo.Read(c.Request.Context(), id)

	if err != nil {
		b.config.Log.Errorf(err.Error())
		c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Read all cards from DB
// @Description This endpoint retrieve all cards from the database
// @ID read-all-cards-from-db
// @Accept  json
// @Produce json
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPClientError
// @Failure 500 {object} models.HTTPBackendError.
// @Router /get-all-cards [get]
func (b *BizcardController) ReadAllBizCard(c *gin.Context) {
	res, err := b.bizcardRepo.ReadAll(c.Request.Context())

	if err != nil {
		b.config.Log.Errorf(err.Error())
		c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	c.JSON(http.StatusOK, res)
}
