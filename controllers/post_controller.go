package controllers

import (
	"fmt"
	"log"
	"net/http"
	"travelgo/models"
	"travelgo/services"

	"github.com/gin-gonic/gin"
)

// PostController handles post-related operations.
type PostController struct {
	postService    *services.PostService
	storageService *services.StorageService
	urlService     *services.UrlService
}

// NewPostController initializes a new PostController.
// It logs an error if the PostService cannot be initialized.
func NewPostController() *PostController {
	postService, err := services.NewPostService()
	if err != nil {
		log.Fatalf("Init Post Service error: %v", err)
	}

	return &PostController{
		postService:    postService,
		storageService: services.NewStorageService(),
	}
}

// UploadFile handles the HTTP POST request to upload a file.
// It extracts the file from the form, uploads it using the StorageService,
// and returns a JSON response with the file URL or an error.
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

// UploadLink handles the HTTP GET request to fetch metadata for a given URL.
// It returns a JSON response with the metadata or an error if fetching fails.
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

// CreatePost handles the HTTP POST request to create a new post.
// It binds the JSON payload to a model and calls the PostService to create the post.
// Returns a JSON response with the post ID or an error.
func (pc *PostController) CreatePost(c *gin.Context) {
	var postContent models.PostUploadContent
	if err := c.ShouldBindJSON(&postContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

// UpdatePost handles the HTTP PUT request to update an existing post.
// It binds the JSON payload to a model and calls the PostService to update the post.
// Returns a JSON response with the post ID or an error.
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

// GetPosts handles the HTTP GET request to retrieve all posts.
// It returns a JSON response with the list of posts or an error if none are found.
func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := pc.postService.GetPosts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost handles the HTTP GET request to retrieve a single post by its ID.
// It returns a JSON response with the post content or an error if the post is not found.
func (pc *PostController) GetPost(c *gin.Context) {
	postId := c.Param("id")
	fmt.Println("Blog Id: ", postId)
	post, err := pc.postService.GetPostContent(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
