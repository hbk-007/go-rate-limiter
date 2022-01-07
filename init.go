package main

import (
	"net/http"
	"rateLimiter/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/hello", middleware.RateLimit, sayHello)
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
