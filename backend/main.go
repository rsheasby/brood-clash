package main

import (
	"time"

	"github.com/gin-contrib/cors"
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

// @BasePath /api/v1

// @securityDefinitions.apikey CodeAuth
// @in header
// @name Authorization

func main() {
	r := gin.New()

	// Global Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders:     []string{"Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// Serve Swagger Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1", middleware.Auth)

	// Authenticated requests go here
	api.GET("/test", controllers.Test)
	api.GET("/currentQuestion", controllers.GetCurrentQuestion)
	api.GET("/questions", controllers.GetAllQuestions)
	api.POST("/questions", controllers.PostQuestions)
	api.DELETE("/questions/:id", controllers.DeleteQuestion)
	api.POST("/questions/:id/select", controllers.SelectQuestion)
	api.POST("/answers/:id/reveal", controllers.RevealAnswer)
	api.POST("/reset", controllers.ResetGameState)
	api.POST("/incorrectAnswer", controllers.IncorrectAnswer)

	// Anonymous requests go here
	r.GET("/presenter/websocket", controllers.GetPresenterWebsocket)

	r.Run(":8080")
}
