package main

import (
	"backend/api"
	"backend/services"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			var err error
			var msg string

			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				continue
			}
			fmt.Printf("%s\n", msg)

			err = websocket.Message.Send(ws, msg)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	g := e.Group("/api")
	g.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:Authorization",
		Validator: services.LoginCodeValidator,
	}))

	g.GET("/code/valid", api.ValidateCode)
	g.GET("/questions", api.GetQuestions)
	g.GET("/questions/:id", api.GetQuestion)
	g.POST("/questions", api.AddQuestions)
	g.PUT("/questions/:questionId/answers/:answerId/revealed", api.RevealAnswer)

	e.GET("/ws", hello)

	e.Start(":1323")
}
