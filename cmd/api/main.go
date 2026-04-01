package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine

	router = gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server Running",
		})
	})

	router.Run(":8080")
}
