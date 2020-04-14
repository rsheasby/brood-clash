package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/controllers"
	_ "github.com/rsheasby/brood-clash/backend/docs"
	"github.com/rsheasby/brood-clash/backend/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Brood-Clash API
// @version 1.0
// @licence.name MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey CodeAuth
// @in header
// @name Authorization

func main() {
	r := gin.New()

	// Global Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Serve Swagger Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	auth := api.Group("", middleware.Auth)
	anon := api.Group("")

	// Authenticated requests go here
	auth.GET("/unshownQuestion", controllers.GetUnshownQuestion)
	auth.GET("/currentQuestion", controllers.GetCurrentQuestion)
	auth.POST("/questions", controllers.PostQuestions)
	auth.POST("/answers/:id/reveal", controllers.RevealAnswer)

	// Anonymous requests go here
	anon.GET("/presenter/websocket", controllers.GetPresenterWebsocket)

	r.Run(":8080")
}
