package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/username", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message" : "hello, jierui303! \n",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
