package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

// AuthController handles HTTP requests related to authentication
type AuthController struct {
	authService services.IAuthService
}

// NewAuthController creates and returns a new AuthController
func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Logout logs out a user
func (c *AuthController) Logout(ctx *gin.Context) {
	// TODO: Implement Logout
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ValidateToken validates an authentication token
func (c *AuthController) ValidateToken(ctx *gin.Context) {
	// TODO: Implement ValidateToken
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// RefreshToken refreshes an authentication token
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	// TODO: Implement RefreshToken
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
