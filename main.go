package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.Default()
	router.Use(AuthMiddleware())

	router.GET("/", Index)

	router.Run(":" + port)
}
