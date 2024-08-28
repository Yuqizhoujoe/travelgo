package controllers

import (
	"net/http"
	"travelgo/models"
	"travelgo/services"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related operations.
type UserController struct {
	userService *services.UserService
}

// NewUserController initializes a new UserController.
// It returns an error if the user service cannot be initialized.
func NewUserController() (*UserController, error) {
	userService, err := services.NewUserService()
	if err != nil {
		return nil, err
	}

	return &UserController{
		userService: userService,
	}, nil
}

// CreateUser handles the HTTP POST request to create a new user.
// It binds the JSON payload to a model and calls the UserService to add the user.
// Returns a JSON response with the result or an error.
func (uc *UserController) CreateUser(c *gin.Context) {
	var addUserModel models.AddUser
	if err := c.ShouldBindJSON(&addUserModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addUserResponse, err := uc.userService.AddUser(addUserModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, addUserResponse)
}
