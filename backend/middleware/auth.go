package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/services"
	"net/http"
)

func Auth(c * gin.Context) {
	if c == nil {
		return
	}

	if services.IsValidCode(c.GetHeader("Authorization")) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
