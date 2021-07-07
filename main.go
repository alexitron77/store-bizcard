package main

import (
	"github.com/gin-gonic/gin"

	ctrl "biz.card/api/controllers"
	repo "biz.card/api/repositories"
	conf "biz.card/config"
)

func main() {
	r := gin.Default()

	db := conf.ConnectDB()
	bizcardRepo := repo.NewBizCardModel(db)
	card := ctrl.NewBizcardController(bizcardRepo)

	r.GET("/save", card.SaveBizCard)

	r.Run()
}
