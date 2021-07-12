package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	ctrl "biz.card/api/controllers"
	"biz.card/api/repositories/mongo"
	"biz.card/config"
)

func main() {
	r := gin.Default()

	path, _ := filepath.Abs("config/env")
	conf := config.LoadConfig(path)

	db := &mongo.DBConn{
		Url:      conf.Mongo.Url,
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
	}

	conn := db.ConnectDB()
	defer conn.DB.Disconnect(conn.Ctx)

	bizcardRepo := mongo.NewBizCardModel(conn.DB, conn.Ctx)
	card := ctrl.NewBizcardController(bizcardRepo)

	r.POST("/create-card", card.SaveBizCard)
	r.POST("/upload-card", card.Upload)

	r.Run()
}
