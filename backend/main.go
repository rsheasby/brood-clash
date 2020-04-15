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

	api := r.Group("/api/v1", middleware.Auth)

	// Authenticated requests go here
	api.GET("/unshownQuestion", controllers.GetUnshownQuestion)
	api.GET("/currentQuestion", controllers.GetCurrentQuestion)
	api.GET("/questions", controllers.GetAllQuestions)
	api.POST("/questions", controllers.PostQuestions)
	api.POST("/answers/:id/reveal", controllers.RevealAnswer)
	api.POST("/reset", controllers.ResetGameState)

	// Anonymous requests go here
	r.GET("/presenter/websocket", controllers.GetPresenterWebsocket)

	r.Run(":8080")
}
