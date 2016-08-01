package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World, Go!"})
}
