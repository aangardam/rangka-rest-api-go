package main

import (
	"fmt"
	"rangka-rest-api-go/config"
	"rangka-rest-api-go/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	portStr := config.ENV.PORT
	port := helper.GetAvalaiblePort(portStr)
	r.Run(fmt.Sprintf(":%d", port))
}
