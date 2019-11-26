package main

import (
	"backend/api"
	ffmiddleware "backend/middleware"
	"backend/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	e.GET("/present", api.Present)

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

	e.HideBanner = true
	_ = e.Start(":1323")
}
