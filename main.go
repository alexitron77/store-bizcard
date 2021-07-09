package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	ctrl "biz.card/api/controllers"
	repo "biz.card/api/repositories"
	"biz.card/config"
	"biz.card/pkg/mongo"
)

func main() {
	r := gin.Default()

	path, _ := filepath.Abs("config/env")
	conf := config.LoadConfig(path)

	mongo := &mongo.MongoModel{
		Url:      conf.Mongo.Url,
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
	}

	db, ctx := mongo.ConnectDB()
	defer db.Disconnect(ctx)

	bizcardRepo := repo.NewBizCardModel(db, ctx)
	card := ctrl.NewBizcardController(bizcardRepo)

	r.GET("/save", card.SaveBizCard)

	r.Run()
}
