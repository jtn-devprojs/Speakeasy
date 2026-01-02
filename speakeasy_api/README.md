# Speakeasy API

A high-performance REST API built with Go for the Speakeasy platform. Designed to serve mobile and web clients with user authentication, profile management, and preference handling.

## Architecture

This project uses the **Gin Web Framework** for HTTP routing and the **Dependency Injection** pattern with a centralized container for managing service dependencies, ensuring loose coupling and testability.

### Folder Structure

```
speakeasy_api/
├── cmd/
│   └── server/
│       └── main.go                      # Server entry point
├── internal/
│   ├── di/
│   │   ├── container.go                 # Dependency injection container
│   │   └── di_test.go                   # DI container unit tests
│   ├── routes/
│   │   └── routes.go                    # Route registration
│   ├── controllers/
│   │   ├── user_controller.go           # User HTTP controllers
│   │   ├── user_controller_test.go      # User controller tests
│   │   ├── auth_controller.go           # Authentication HTTP controllers
│   │   ├── auth_controller_test.go      # Auth controller tests
│   │   ├── location_controller.go       # Location HTTP controllers
│   │   └── location_controller_test.go  # Location controller tests
│   └── services/
│       ├── user_service.go              # User business logic
│       ├── user_service_test.go         # User service unit tests
│       ├── auth_service.go              # Auth business logic
│       ├── auth_service_test.go         # Auth service unit tests
│       ├── location_service.go          # Location business logic
│       ├── location_service_test.go     # Location service unit tests
│       └── errors.go                    # Service error definitions
├── go.mod                               # Go module definition
└── README.md                            # This file
```

## Testing

Unit tests are co-located with the code they test (white-box testing approach):
- `internal/services/user_service_test.go` - UserService tests
- `internal/services/auth_service_test.go` - AuthService tests
- `internal/di/di_test.go` - DI container tests

Run tests with:
```bash
go test ./...
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
// Access services and controllers
userService := container.UserService
authService := container.AuthService
userController := container.UserController
authController := container.AuthController
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

### Controller Implementation
- [ ] UserController.GetUser
- [ ] UserController.CreateUser
- [ ] UserController.UpdateUser
- [ ] UserController.DeleteUser
- [ ] UserController.GetUserPreferences
- [ ] UserController.UpdateUserPreferences
- [ ] AuthController.Login
- [ ] AuthController.Logout
- [ ] AuthController.Register
- [ ] AuthController.ValidateToken
- [ ] AuthController.RefreshToken

### Service Implementation
- [ ] User service methods
- [ ] Auth service methods
- [ ] Database integration
- [ ] Token management
- [ ] Error handling

## Project Structure Details

```
internal/
├── di/
│   └── container.go          # DI container - manages all dependencies
├── routes/
│   └── routes.go             # Route registration and grouping
├── controllers/
│   ├── user_controller.go    # HTTP controllers for user operations
│   ├── auth_controller.go    # HTTP controllers for auth operations
│   └── location_controller.go # HTTP controllers for location operations
└── services/
    ├── user_service.go       # User business logic
    ├── auth_service.go       # Auth business logic
    ├── location_service.go   # Location business logic
    └── errors.go             # Custom error definitions
```

## Next Steps

1. Implement database layer (repository pattern)
2. Add authentication (JWT tokens)
3. Implement request validation
4. Add middleware for logging and error handling
5. Configure environment variables
6. Add API documentation (OpenAPI/Swagger)
