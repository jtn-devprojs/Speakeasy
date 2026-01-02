package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/speakeasy/speakeasy-api/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Login
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Logout
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Register
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement ValidateToken
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement RefreshToken
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}
