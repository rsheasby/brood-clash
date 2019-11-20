package main

import (
	"backend/api"
	"backend/services"

	"fmt"
	"net/http"

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

// TODO: Need a better name for the function
func foo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		switch c.Request().Method {
		case http.MethodPost:
		case http.MethodPut:
		case http.MethodPatch:
		case http.MethodDelete:
		default:
			return next(c)
		}

		c.Response().Before(func() {
			if c.Response().Status/100 == 2 {
				println("Ladies and gentlemen...")
			}
		})
		return next(c)
	}
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
	q.Use(foo)
	q.GET("", api.GetQuestions)
	q.POST("", api.AddQuestions)
	q.GET("/:id", api.GetQuestion)
	q.PUT("/:questionId/answers/:answerId/revealed", api.RevealAnswer)

	// e.GET("/present", api.Present)

	e.Start(":1323")
}
