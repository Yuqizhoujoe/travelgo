package main

import (
	"log"
	"travelgo/firebase"

	"travelgo/controllers"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// Env Config Init
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	firebase.InitFirebase()

	r := gin.Default()

	postController := controllers.NewPostController()

	r.POST("/travel/posts/upload", postController.UploadFile)
	r.POST("/travel/posts", postController.CreatePost)
	r.GET("/travel/posts", postController.GetPosts)
	r.GET("/travel/posts/:id", postController.GetPost)

	r.Run(":8081")
}
