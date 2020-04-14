package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/services"
)

func Auth(c *gin.Context) {
	if c == nil {
		return
	}

	if services.IsValidCode(c.GetHeader("Authorization")) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
