package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Store-Token")
		secretToken := os.Getenv("SECRET_TOKEN")

		if token != secretToken {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
