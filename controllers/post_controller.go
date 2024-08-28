package controllers

import (
	"fmt"
	"net/http"
	"travelgo/models"
	"travelgo/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService    services.PostService
	storageService services.StorageService
	urlService     services.UrlService
}

func NewPostController() *PostController {
	return &PostController{
		postService:    services.NewPostService(),
		storageService: services.NewStorageService(),
	}
}

func (pc *PostController) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := pc.storageService.UploadFile(file, header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": link})
}

func (pc *PostController) UploadLink(c *gin.Context) {
	urlParams := c.Query("url")
	if urlParams == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	metadata, err := pc.urlService.FetchMetadata(urlParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch URL metadata"})
		return
	}

	response := models.FetchUrlResponse{
		Success: 1,
		Meta:    *metadata,
	}

	c.JSON(http.StatusOK, response)
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var postContent models.PostUploadContent
	if err := c.ShouldBindJSON(&postContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // Handle post thumbnail upload
	// file, header, err := c.Request.FormFile("postThumbnail")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to uplaod thumbnail"})
	// 	return
	// }

	// // Call storage service to upload file and get the URL
	// gcsURL, err := pc.storageService.UploadFile(file, header)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Call postService.CreatePost with uploaded file
	fmt.Println("Post Controller")
	fmt.Println(postContent)
	postId, err := pc.postService.CreatePost(postContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.PostResponse{
		Success: true,
		PostID:  postId,
	}

	c.JSON(http.StatusCreated, response)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	postId := c.Param("id")
	var postContent models.PostUploadContent
	if err := c.ShouldBindBodyWithJSON(&postContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId, err := pc.postService.UpdatePost(postId, postContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.PostResponse{
		Success: true,
		PostID:  postId,
	}

	c.JSON(http.StatusCreated, response)
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
	postId := c.Param("id")
	post, err := pc.postService.GetPostContent(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
