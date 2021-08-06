package api

import (
	ctrl "biz.card/api/controllers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Handlers(r *gin.Engine, card *ctrl.BizcardController) {
	r.POST("/create-card", card.SaveBizCard, card.Upload, card.UpdateCardURL)
	r.GET("/get-card/:id", card.ReadBizCard)
	r.GET("/get-all-cards", card.ReadAllBizCard)
	r.GET("/ws", card.ConnWebSocket)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
