# Speakeasy API

A high-performance REST API built with Go for the Speakeasy platform. Designed to validate Firebase authentication tokens and manage user sessions and locations.

## Architecture

This project uses:
- **Gin Web Framework** for HTTP routing
- **Dependency Injection (DI) Pattern** with a centralized container for managing service dependencies
- **Database Abstraction** through interfaces to support multiple databases (PostgreSQL, SQLite)
- **Interface-based Locking** (`ISessionLocker`) for database-specific pessimistic locking

### Key Design Decisions

1. **No User Management:** Firebase handles user authentication on the client side. The server only validates tokens.
2. **Token-Based Auth:** All endpoints require Firebase ID tokens in the Authorization header.
3. **Database Abstraction:** Uses interfaces to abstract database operations, allowing seamless switching between PostgreSQL (production) and SQLite (testing).
4. **Locking Strategy:**
   - **PostgreSQL:** Uses `FOR UPDATE` for row-level pessimistic locking
   - **SQLite:** Uses implicit transaction locking via SELECT queries

### Folder Structure

```
speakeasy_api/
├── cmd/
│   └── server/
│       └── main.go                      # Server entry point
├── internal/
│   ├── config/
│   │   └── config.go                    # Configuration management
│   ├── database/
│   │   └── database.go                  # Database connection setup
│   ├── di/
│   │   ├── container.go                 # DI container with dbType switch for locker selection
│   │   └── di_test.go                   # DI container unit tests
│   ├── routes/
│   │   └── routes.go                    # Route registration
│   ├── controllers/
│   │   ├── auth_controller.go           # Token validation endpoints (Logout, ValidateToken, RefreshToken)
│   │   ├── auth_controller_test.go      # Auth controller tests
│   │   ├── session_controller.go        # Session & location management endpoints
│   │   └── session_controller_test.go   # Session controller tests
│   ├── services/
│   │   ├── interfaces.go                # Service interfaces (IAuthService, ISessionService)
│   │   ├── auth_service.go              # Token validation logic
│   │   ├── auth_service_test.go         # Auth service unit tests
│   │   ├── session_service.go           # Session & location business logic
│   │   ├── session_service_test.go      # Session service unit tests
│   │   └── errors.go                    # Service error definitions
│   └── repositories/
│       ├── interfaces.go                # Repository interfaces (ISessionLocker, ISessionUserRepository, etc.)
│       ├── session_user_repository.go   # Session user data operations with ISessionLocker injection
│       ├── session_user_repository_test.go # Session user tests
│       ├── session_repository.go        # Session data operations
│       ├── user_repository.go           # User data operations
│       └── ...
├── go.mod                               # Go module definition
└── README.md                            # This file
```

## API Endpoints

### Authentication
- `POST /api/auth/logout` - Logout user (revoke token)
- `POST /api/auth/validate` - Validate Firebase token
- `POST /api/auth/refresh` - Refresh authentication token

### Sessions & Locations
- `POST /api/sessions/check-vicinity` - Check if user is near a session
- `GET /api/sessions/nearby` - Get nearby sessions
- `GET /api/sessions/location` - Get user's current location
- `PUT /api/sessions/location` - Update user's location

## Database Abstraction

### ISessionLocker Interface
Provides database-specific locking implementations:

```go
type ISessionLocker interface {
    LockSession(ctx context.Context, tx interface{}, sessionID string) error
}
```

**Implementations:**
- `PostgresSessionLocker` - Uses `FOR UPDATE` for row-level locking
- `SqliteSessionLocker` - Uses implicit transaction locking

The DI container selects the appropriate locker based on `dbType`:

```go
switch dbType {
case "postgres":
    locker = &repositories.PostgresSessionLocker{}
case "sqlite":
    locker = &repositories.SqliteSessionLocker{}
}
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
│   └── session_controller.go # HTTP controllers for session operations
└── services/
    ├── user_service.go       # User business logic
    ├── auth_service.go       # Auth business logic
    ├── session_service.go    # Session business logic
    └── errors.go             # Custom error definitions
```

## Next Steps

1. Implement database layer (repository pattern)
2. Add authentication (JWT tokens)
3. Implement request validation
4. Add middleware for logging and error handling
5. Configure environment variables
6. Add API documentation (OpenAPI/Swagger)
