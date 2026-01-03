package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/di"
	"github.com/speakeasy/speakeasy-api/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, container *di.Container) {
	// Apply auth middleware to protected routes
	sessionGroup := router.Group("/api/sessions")
	sessionGroup.Use(middleware.AuthMiddleware(container.AuthService))
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
