package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

// UserController handles HTTP requests related to users
type UserController struct {
	userService services.IUserService
}

// NewUserController creates and returns a new UserController
func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetUser retrieves a user by ID
func (c *UserController) GetUser(ctx *gin.Context) {
	// TODO: Implement GetUser
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateUser creates a new user
func (c *UserController) CreateUser(ctx *gin.Context) {
	// TODO: Implement CreateUser
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateUser updates an existing user
func (c *UserController) UpdateUser(ctx *gin.Context) {
	// TODO: Implement UpdateUser
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteUser deletes a user
func (c *UserController) DeleteUser(ctx *gin.Context) {
	// TODO: Implement DeleteUser
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetUserPreferences retrieves user preferences
func (c *UserController) GetUserPreferences(ctx *gin.Context) {
	// TODO: Implement GetUserPreferences
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateUserPreferences updates user preferences
func (c *UserController) UpdateUserPreferences(ctx *gin.Context) {
	// TODO: Implement UpdateUserPreferences
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
