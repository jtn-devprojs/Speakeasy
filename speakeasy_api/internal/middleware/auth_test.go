package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	mockAuthService := &services.MockAuthService{
		ValidateTokenFunc: func(token string) (string, error) {
			if token == "valid-token" {
				return "user123", nil
			}
			return "", services.ErrNotImplemented
		},
	}
	middleware := AuthMiddleware(mockAuthService)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/sessions", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	handlerCalled := false
	router := gin.New()
	router.GET("/api/sessions", middleware, func(c *gin.Context) {
		handlerCalled = true
		userID, exists := c.Get("userID")
		if !exists {
			t.Fatal("Expected userID to be set in context")
		}
		if userID != "user123" {
			t.Fatalf("Expected userID to be 'user123', got %v", userID)
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.ServeHTTP(w, req)

	if !handlerCalled {
		t.Fatal("Expected handler to be called")
	}
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}

func TestAuthMiddleware_MissingAuthHeader(t *testing.T) {
	mockAuthService := &services.MockAuthService{}
	middleware := AuthMiddleware(mockAuthService)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/sessions", nil)
	// No Authorization header

	router := gin.New()
	router.GET("/api/sessions", middleware, func(c *gin.Context) {
		t.Fatal("Expected handler not to be called")
	})

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidHeaderFormat(t *testing.T) {
	mockAuthService := &services.MockAuthService{}
	middleware := AuthMiddleware(mockAuthService)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/sessions", nil)
	req.Header.Set("Authorization", "InvalidFormat token")

	router := gin.New()
	router.GET("/api/sessions", middleware, func(c *gin.Context) {
		t.Fatal("Expected handler not to be called")
	})

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	mockAuthService := &services.MockAuthService{
		ValidateTokenFunc: func(token string) (string, error) {
			return "", services.ErrNotImplemented
		},
	}
	middleware := AuthMiddleware(mockAuthService)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/sessions", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	router := gin.New()
	router.GET("/api/sessions", middleware, func(c *gin.Context) {
		t.Fatal("Expected handler not to be called")
	})

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_BearerWithoutToken(t *testing.T) {
	mockAuthService := &services.MockAuthService{}
	middleware := AuthMiddleware(mockAuthService)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/sessions", nil)
	req.Header.Set("Authorization", "Bearer")

	router := gin.New()
	router.GET("/api/sessions", middleware, func(c *gin.Context) {
		t.Fatal("Expected handler not to be called")
	})

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_BearerWithEmptyToken(t *testing.T) {
	mockAuthService := &services.MockAuthService{}
	middleware := AuthMiddleware(mockAuthService)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/sessions", nil)
	req.Header.Set("Authorization", "Bearer ")

	router := gin.New()
	router.GET("/api/sessions", middleware, func(c *gin.Context) {
		t.Fatal("Expected handler not to be called")
	})

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status 401, got %d", w.Code)
	}
}
