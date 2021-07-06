package main

import (
	ct "biz.card/api/controllers"
	"github.com/gin-gonic/gin"

	rp "biz.card/api/repositories"
	cf "biz.card/config"
)

func main() {
	r := gin.Default()

	db := cf.ConnectDB()

	bizcardRepo := rp.NewBizCardRepo(db)

	h := ct.NewBizcardHandler(bizcardRepo)

	r.GET("/health", h.SaveBizCard)

	r.Run()
}
