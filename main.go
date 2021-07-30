package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	ctrl "biz.card/api/controllers"
	"biz.card/api/repositories/aws"
	"biz.card/api/repositories/mongo"
	"biz.card/config"
	_ "biz.card/docs"
	mw "biz.card/middleware"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @basePath /
func main() {
	path, _ := filepath.Abs("config/env")
	conf := config.LoadConfig(path)

	r := gin.Default()
	r.Use(mw.GinLogMiddleware())

	s3 := aws.AwsInit(conf.Aws.AccessKey, conf.Aws.Secret)

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})

	db := &mongo.DBConn{
		Url:      conf.Mongo.Url,
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
	}

	dbconn := db.ConnectDB()
	defer dbconn.DB.Disconnect(dbconn.Ctx)

	config := config.NewConfig(dbconn.DB, dbconn.Ctx, logger, s3)
	bizCardRepo := mongo.NewDBRepo(config)
	awsRepo := aws.NewAwsRepo()

	card := ctrl.NewBizcardController(config, bizCardRepo, awsRepo)

	// Handlers chain on create-card endpoint
	r.POST("/create-card", card.SaveBizCard, card.Upload, card.UpdateCardURL)
	r.GET("/get-card/:name", card.ReadBizCard)
	r.GET("/ws", card.ConnWebSocket)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
