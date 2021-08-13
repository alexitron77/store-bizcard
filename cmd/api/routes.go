package api

import (
	"crypto/tls"
	"fmt"
	"net/http"

	ctrl "biz.card/cmd/api/controllers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type api struct {
	addr *string
	r    *gin.Engine
}

func NewApi(addr *string, r *gin.Engine) *api {
	return &api{
		addr,
		r,
	}
}

func (a *api) Routes(card *ctrl.BizcardController) http.Handler {
	a.r.POST("/create-card", card.SaveBizCard, card.Upload, card.UpdateCardURL)
	a.r.GET("/get-card/:id", card.ReadBizCard)
	a.r.GET("/get-all-cards", card.ReadAllBizCard)
	a.r.GET("/ws", card.ConnWebSocket)
	a.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return a.r
}

func (a *api) Start() {
	server := http.Server{
		Addr:      *a.addr,
		Handler:   a.r,
		TLSConfig: &tls.Config{},
	}

	err := server.ListenAndServeTLS("./certificate/cert.pem", "./certificate/key.pem")

	if err != nil {
		fmt.Print(err)
	}
}
