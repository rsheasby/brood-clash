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
	auth.GET("/question", controllers.GetUnshownQuestion)
	auth.POST("/questions", controllers.PostQuestions)

	// Anonymous requests go here
	anon.GET("/presenter/websocket",  controllers.GetPresenterWebsocket)

	r.Run(":8080")
}
