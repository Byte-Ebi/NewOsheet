package main

import (
	"NewOsheet/internal/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Middleware())

	r.GET("/hc", healthCheck)

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":80")
	if err != nil {
		return
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
