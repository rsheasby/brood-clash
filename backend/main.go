package main

import (
	"backend/api"
	ffmiddleware "backend/middleware"
	"backend/models"
	"backend/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		ch := make(chan []models.Question)
		services.RegisterUpdateListener(ch)
		defer services.UnregisterUpdateListener(ch)

		for {
			err := websocket.JSON.Send(ws, <-ch)
			if err != nil {
				c.Logger().Error(err)
				break
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

	q := g.Group("/questions")
	q.Use(ffmiddleware.DetectQuestionUpdates)
	q.GET("", api.GetQuestions)
	q.POST("", api.AddQuestions)
	q.GET("/:id", api.GetQuestion)
	q.PUT("/:questionId/answers/:answerId/revealed", api.RevealAnswer)

	e.GET("/present", hello)

	e.Start(":1323")
}
