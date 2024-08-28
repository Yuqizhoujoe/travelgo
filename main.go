package main

import (
	"log"
	"time"
	"travelgo/firebase"

	"travelgo/controllers"

	"github.com/gin-contrib/cors"
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

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	postController := controllers.NewPostController()

	r.POST("/travel/posts/upload-image", postController.UploadFile)
	r.POST("/travel/posts", postController.CreatePost)

	r.PUT("/travel/posts/:id", postController.UpdatePost)

	r.GET("/travel/posts/upload-link", postController.UploadLink)
	r.GET("/travel/posts", postController.GetPosts)
	r.GET("/travel/posts/:id", postController.GetPost)

	r.Run(":8081")
}
