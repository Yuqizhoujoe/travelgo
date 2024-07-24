package controllers

import (
	"net/http"
	"travelgo/models"
	"travelgo/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService    services.PostService
	storageService services.StorageService
}

func NewPostController() *PostController {
	return &PostController{
		postService:    services.NewPostService(),
		storageService: services.NewStorageService(),
	}
}

func (pc *PostController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := pc.storageService.UploadFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"link": link})
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var postContent models.PostUploadContent
	if err := c.ShouldBindJSON(&postContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.postService.CreatePost(postContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := pc.postService.GetPosts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPost(c *gin.Context) {
	postID := c.Param("id")
	post, err := pc.postService.GetPostContent(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
