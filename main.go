package main

import (
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test is passed!",
		})
	})

	log.Println("server start at port 8080")
	r.Run(":8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}

