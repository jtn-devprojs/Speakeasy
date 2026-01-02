package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/speakeasy/speakeasy-api/internal/services"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement GetUser
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement CreateUser
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement UpdateUser
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement DeleteUser
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *UserHandler) GetUserPreferences(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement GetUserPreferences
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}

func (h *UserHandler) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement UpdateUserPreferences
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}
