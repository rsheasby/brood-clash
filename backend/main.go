package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/controllers"
	"github.com/rsheasby/brood-clash/backend/middleware"
	"github.com/rsheasby/brood-clash/backend/models"
	"github.com/rsheasby/brood-clash/backend/services"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/rsheasby/brood-clash/backend/docs"
)

// @title Swagger Example API
// @version 1.0
// @licence.name Apache 2.0

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
	auth.GET("/authPing", func(c *gin.Context) {
		// This was just for testing the database stuffs.
		// TODO: Remember to remove
		err := services.InsertQuestion(models.Question{
			Text: "test",
			Answers: []models.Answer{
				{
					Text:       "testanswer",
					Points:     1,
					Revealed:   false,
				},
			},
		})
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to add question: %v", err))
		}
		c.Status(200)
	})

	// Anonymous requests go here
	anon.GET("/ping", controllers.Pong)

	r.Run(":8080")
}
