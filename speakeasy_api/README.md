# Speakeasy API

A high-performance REST API built with Go for the Speakeasy platform. Designed to serve mobile and web clients with user authentication, profile management, and preference handling.

## Architecture

This project follows the **Dependency Injection** pattern with a centralized container for managing service dependencies, ensuring loose coupling and testability.

### Folder Structure

```
speakeasy_api/
├── cmd/
│   └── server/
│       └── main.go              # Server entry point
├── internal/
│   ├── di/
│   │   └── container.go         # Dependency injection container
│   ├── handlers/
│   │   ├── user_handler.go      # User HTTP handlers
│   │   └── auth_handler.go      # Authentication HTTP handlers
│   ├── services/
│   │   ├── user_service.go      # User business logic
│   │   ├── auth_service.go      # Auth business logic
│   │   └── errors.go            # Service error definitions
│   └── ...
├── test/
│   ├── di_test.go               # DI container tests
│   ├── user_service_test.go     # User service tests
│   └── auth_service_test.go     # Auth service tests
├── go.mod                        # Go module definition
└── README.md                     # This file
```

## Services

### UserService
Manages user-related business logic.
- `GetUserByID()` - Retrieve user by ID
- `CreateUser()` - Create new user
- `UpdateUser()` - Update user information
- `DeleteUser()` - Delete user account
- `GetUserPreferences()` - Retrieve user preferences
- `UpdateUserPreferences()` - Update user preferences

### AuthService
Handles authentication and authorization.
- `Login()` - Authenticate user and return token
- `Logout()` - Invalidate user session
- `Register()` - Create new user account
- `ValidateToken()` - Verify authentication token
- `RefreshToken()` - Generate new token

## API Endpoints

### Authentication
- `POST /api/auth/login` - User login
- `POST /api/auth/logout` - User logout
- `POST /api/auth/register` - User registration
- `POST /api/auth/validate` - Token validation
- `POST /api/auth/refresh` - Token refresh

### Users
- `GET /api/users/{id}` - Get user
- `POST /api/users` - Create user
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user
- `GET /api/users/{id}/preferences` - Get user preferences
- `PUT /api/users/{id}/preferences` - Update user preferences

### Health Check
- `GET /api/health` - Server health status

## Dependency Injection

The DI container is initialized in `internal/di/container.go`:

```go
container := di.NewContainer()
// Access services
userService := container.UserService
authService := container.AuthService
userHandler := container.UserHandler
authHandler := container.AuthHandler
```

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Installation

1. Navigate to the project directory:
```bash
cd speakeasy_api
```

2. Download dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

Run specific test package:
```bash
go test ./test
```

## Implementation Status

All handler and service methods are currently stubbed with `TODO` markers. Implementation will follow in subsequent phases.

### Handler Implementation
- [ ] UserHandler.GetUser
- [ ] UserHandler.CreateUser
- [ ] UserHandler.UpdateUser
- [ ] UserHandler.DeleteUser
- [ ] UserHandler.GetUserPreferences
- [ ] UserHandler.UpdateUserPreferences
- [ ] AuthHandler.Login
- [ ] AuthHandler.Logout
- [ ] AuthHandler.Register
- [ ] AuthHandler.ValidateToken
- [ ] AuthHandler.RefreshToken

### Service Implementation
- [ ] User service methods
- [ ] Auth service methods
- [ ] Database integration
- [ ] Token management
- [ ] Error handling

## Project Structure

```
internal/
├── di/
│   └── container.go          # DI container - manages all dependencies
├── handlers/
│   ├── user_handler.go       # HTTP handlers for user operations
│   └── auth_handler.go       # HTTP handlers for auth operations
└── services/
    ├── user_service.go       # User business logic
    ├── auth_service.go       # Auth business logic
    └── errors.go             # Custom error definitions
```

## Next Steps

1. Implement database layer (repository pattern)
2. Add authentication (JWT tokens)
3. Implement request validation
4. Add middleware for logging and error handling
5. Configure environment variables
6. Add API documentation (OpenAPI/Swagger)
