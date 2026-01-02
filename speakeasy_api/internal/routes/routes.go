package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/di"
)

func RegisterRoutes(router *gin.Engine, container *di.Container) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", container.AuthController.Login)
		authGroup.POST("/logout", container.AuthController.Logout)
		authGroup.POST("/register", container.AuthController.Register)
		authGroup.POST("/validate", container.AuthController.ValidateToken)
		authGroup.POST("/refresh", container.AuthController.RefreshToken)
	}

	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/:id", container.UserController.GetUser)
		userGroup.POST("", container.UserController.CreateUser)
		userGroup.PUT("/:id", container.UserController.UpdateUser)
		userGroup.DELETE("/:id", container.UserController.DeleteUser)
		userGroup.GET("/:id/preferences", container.UserController.GetUserPreferences)
		userGroup.PUT("/:id/preferences", container.UserController.UpdateUserPreferences)
	}

	locationGroup := router.Group("/api/locations")
	{
		locationGroup.POST("/check-vicinity", container.LocationController.CheckVicinity)
		locationGroup.GET("/nearby", container.LocationController.GetNearbyLocations)
		locationGroup.GET("/user", container.LocationController.GetUserLocation)
		locationGroup.PUT("/user", container.LocationController.UpdateUserLocation)
	}

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
