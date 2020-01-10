package controllers

import (
	"backend/models"
	"backend/restapi/operations"
	"backend/services"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/net/websocket"
)

func init() {
	ControllerConfigList.Register(func(api *operations.BroodClashAPI) {
		api.WebsocketHandler = operations.WebsocketHandlerFunc(WebsocketHandler)
	})
}

func WebsocketHandler(params operations.WebsocketParams) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		websocket.Handler(func(ws *websocket.Conn) {
			var err error

			defer ws.Close()
			ch := make(chan []models.Question)
			services.RegisterUpdateListener(ch)
			defer services.UnregisterUpdateListener(ch)

			err = websocket.JSON.Send(ws, services.GetQuestions())
			if err != nil {
				// TODO: Log this error somewhere, somehow
				return
			}

			for {
				err = websocket.JSON.Send(ws, <-ch)
				if err != nil {
					// TODO: Log this error.
					break
				}
			}
		}).ServeHTTP(rw, params.HTTPRequest)
	})
}
