package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/di"
)

func RegisterRoutes(router *gin.Engine, container *di.Container) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/logout", container.AuthController.Logout)
		authGroup.POST("/validate", container.AuthController.ValidateToken)
		authGroup.POST("/refresh", container.AuthController.RefreshToken)
	}

	sessionGroup := router.Group("/api/sessions")
	{
		sessionGroup.POST("/check-vicinity", container.SessionController.CheckVicinity)
		sessionGroup.GET("/nearby", container.SessionController.GetNearbyLocations)
		sessionGroup.GET("/location", container.SessionController.GetUserLocation)
		sessionGroup.PUT("/location", container.SessionController.UpdateUserLocation)
	}

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
