package main

import (
	"backend/api"
	"backend/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:Authorization",
		Validator: services.LoginCodeValidator,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/code/valid", api.ValidateCode)
	e.GET("/questions", api.GetQuestions)
	e.POST("/questions", api.AddQuestions)

	e.Start(":1323")
}
