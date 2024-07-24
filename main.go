package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/travel/posts/upload")
	r.POST("/travel/posts")
	r.GET("/travel/posts")
	r.GET("/travel/posts/:id")

	r.Run(":8080")
}
