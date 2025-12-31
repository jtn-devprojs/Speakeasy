package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/speakeasy/speakeasy-api/internal/di"
)

func main() {
	// Initialize dependency injection container
	container := di.NewContainer()

	// Create router
	router := mux.NewRouter()

	// Register auth routes
	router.HandleFunc("/api/auth/login", container.AuthHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/logout", container.AuthHandler.Logout).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/register", container.AuthHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/validate", container.AuthHandler.ValidateToken).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/refresh", container.AuthHandler.RefreshToken).Methods(http.MethodPost)

	// Register user routes
	router.HandleFunc("/api/users/{id}", container.UserHandler.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/users", container.UserHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}", container.UserHandler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}", container.UserHandler.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/{id}/preferences", container.UserHandler.GetUserPreferences).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}/preferences", container.UserHandler.UpdateUserPreferences).Methods(http.MethodPut)

	// Health check endpoint
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	// Start server
	log.Println("Starting Speakeasy API server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
