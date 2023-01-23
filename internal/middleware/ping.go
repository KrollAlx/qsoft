package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func PingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.ToLower(c.GetHeader("X-PING")) == "ping" {
			c.Header("X-PONG", "pong")
		}
		c.Next()
	}
}
