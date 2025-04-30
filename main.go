package main

import (
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "EKS Deploy Is Success!!!",
		})
	})

	log.Println("server start at port 8080")
	r.Run(":8080")
}

