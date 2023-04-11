package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Andromeda",
			"bio":  "Software Engineer",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title":    "Hello World",
			"subtitle": "Belajar Gin dan Gorm",
		})
	})

	router.Run(":8888") // Using port :8080 by default
}
