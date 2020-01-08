package controllers

import (
	"backend/restapi/operations"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/net/websocket"
)

func ConfigurePresenterAPI(api *operations.BroodClashAPI) {
	api.WebsocketHandler = operations.WebsocketHandlerFunc(func(params operations.WebsocketParams) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
			websocket.Handler(func(ws *websocket.Conn) {
				var err error

				defer ws.Close()
				// ch := make(chan []models.Question)
				// services.RegisterUpdateListener(ch)
				// defer services.UnregisterUpdateListener(ch)

				err = websocket.Message.Send(ws, "test")
				if err != nil {
					// TODO: problem...
					return
				}

				ch := time.NewTicker(2 * time.Second)

				for {
					<-ch.C
					err = websocket.Message.Send(ws, "yeet")
					if err != nil {
						// TODO: problem....
						break
					}
				}
			}).ServeHTTP(rw, params.HTTPRequest)
		})
	})
}
