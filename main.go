package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	ctrl "biz.card/api/controllers"
	"biz.card/api/repositories/mongo"
	"biz.card/config"
	mw "biz.card/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.Use(mw.GinLogMiddleware())

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})

	path, _ := filepath.Abs("config/env")
	conf := config.LoadConfig(path)

	db := &mongo.DBConn{
		Url:      conf.Mongo.Url,
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
	}

	conn := db.ConnectDB()
	defer conn.DB.Disconnect(conn.Ctx)

	bizcardRepo := mongo.NewBizCardModel(conn.DB, conn.Ctx, logger)
	card := ctrl.NewBizcardController(bizcardRepo, logger)

	r.POST("/create-card", card.SaveBizCard)
	r.POST("/upload-card", card.Upload)
	r.GET("/ws", card.ConnWebSocket)
	r.Run()
}
