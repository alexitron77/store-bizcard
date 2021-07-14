package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var ws *websocket.Conn

// Websocket godoc
// @Summary Create a websocket connection
// @Description This endpoint establish a websocket connection with the client
// @ID connect-websocket
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPClientError
// @Failure 500 {object} models.HTTPBackendError
// @Router /ws [get]
func (b *BizcardController) ConnWebSocket(c *gin.Context) {
	// Allow all origins to avoid CORS issues
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, _ = upgrader.Upgrade(c.Writer, c.Request, nil)

	_, msg, err := ws.ReadMessage()

	if err != nil {
		b.config.Log.Errorf(err.Error())
	}

	fmt.Print(string(msg))
}

func WriteToWs(msg string) {
	err := ws.WriteMessage(1, []byte(msg))

	if err != nil {
		fmt.Print(err)
	}
}
