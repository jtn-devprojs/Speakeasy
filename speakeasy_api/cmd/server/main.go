package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/di"
	"github.com/speakeasy/speakeasy-api/internal/routes"
)

func main() {
	// Initialize dependency injection container
	container := di.NewContainer()

	// Create router
	router := gin.Default()

	// Register all routes
	routes.RegisterRoutes(router, container)

	// Start server
	log.Println("Starting Speakeasy API server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
