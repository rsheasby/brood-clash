package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/controllers"
	"github.com/rsheasby/brood-clash/backend/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

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
	auth.GET("/authPing", func(c *gin.Context) {
		c.Status(200)
	})

	anon := api.Group("")

	anon.GET("/ping", controllers.Pong)

	r.Run(":8080")
}
