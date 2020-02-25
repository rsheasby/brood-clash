package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/rsheasby/brood-clash/backend/docs"
)

// @title Swagger Example API
// @version 1.0
// @licence.name Apache 2.0

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", controllers.Pong)
	r.Run(":8080")
}