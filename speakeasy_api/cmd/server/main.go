package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/config"
	"github.com/speakeasy/speakeasy-api/internal/database"
	"github.com/speakeasy/speakeasy-api/internal/di"
	"github.com/speakeasy/speakeasy-api/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Init(cfg.Database.Type, cfg.Database.Connection)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	container := di.NewContainer(db)

	router := gin.Default()

	routes.RegisterRoutes(router, container)

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Starting Speakeasy API server on port %d (environment: %s)", cfg.Server.Port, cfg.Server.Env)
	if err := router.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
