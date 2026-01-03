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
2. **Middleware-Based Auth:** Token validation is performed via `AuthMiddleware` applied to protected routes. All session endpoints require valid Firebase ID tokens.
3. **Heartbeat-Based Logout:** Users are automatically logged out after inactivity. No explicit logout endpoint needed.
4. **Database Abstraction:** Uses interfaces to abstract database operations, allowing seamless switching between PostgreSQL (production) and SQLite (testing).
5. **Locking Strategy:**
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
    │   └── routes.go                    # Route registration with middleware
    ├── middleware/
    │   ├── auth.go                      # Firebase token validation middleware
    │   └── auth_test.go                 # Middleware unit tests
    ├── controllers/
│   │   ├── session_controller.go        # Session & location management endpoints
│   │   └── session_controller_test.go   # Session controller tests
│   ├── services/
│   │   ├── interfaces.go                # Service interfaces (IAuthService, ISessionService)
│   │   ├── auth_service.go              # Token validation logic
│   │   ├── auth_service_test.go         # Auth service unit tests
│   │   ├── session_service.go           # Session & location business logic
│   │   ├── session_service_test.go      # Session service unit tests
│   │   ├── mocks.go                     # Service mocks for testing
│   │   └── errors.go                    # Service error definitions
│   └── repositories/
│       ├── interfaces.go                # Repository interfaces (ISessionLocker, ISessionUserRepository, etc.)
│       ├── mocks.go                     # Repository mocks for testing
│       ├── session_user_repository.go   # Session user data operations with ISessionLocker injection
    │   ├── session_user_repository_test.go # Session user tests
    │   ├── session_repository.go        # Session data operations
    │   ├── session_repository_test.go   # Session repository tests
    │   ├── user_repository.go           # User data operations
    │   ├── user_repository_test.go      # User repository tests
    │   ├── message_repository.go        # Message data operations
    │   └── message_repository_test.go   # Message repository tests
├── go.mod                               # Go module definition
└── README.md                            # This file
```

## API Endpoints

### Sessions & Locations
All endpoints require Firebase ID token in `Authorization: Bearer <token>` header.

- `POST /api/sessions/check-vicinity` - Check if user is near a session
- `GET /api/sessions/nearby` - Get nearby sessions
- `GET /api/sessions/location` - Get user's current location
- `PUT /api/sessions/location` - Update user's location (resets heartbeat)

### Health Check
- `GET /api/health` - Health check (no auth required)

## Authentication Flow

### Token Validation Middleware
The `AuthMiddleware` is applied to all protected routes:

1. Client sends request with `Authorization: Bearer <firebase-id-token>` header
2. Middleware extracts the token from the header
3. Firebase Admin SDK verifies the token signature and expiration
4. User ID from the token is stored in the request context
5. Handler processes the request with authenticated user
6. If token is invalid or missing, returns 401 Unauthorized

### Logout/Inactivity
Users are automatically removed from sessions after inactivity:

1. Location update (`PUT /api/sessions/location`) resets the user's heartbeat
2. Background cleanup process checks for stale heartbeats (timeout configurable)
3. Users with expired heartbeats are automatically removed from sessions
4. When a session becomes empty, it's marked as ended

No explicit logout endpoint is needed — the app simply stops sending location updates.

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

## Dependency Injection

The DI container is initialized in `internal/di/container.go` and accepts a `dbType` parameter to select database-specific implementations:

```go
container := di.NewContainer(db, "sqlite")  // or "postgres"
// Access services and controllers
authService := container.AuthService
sessionService := container.SessionService
sessionController := container.SessionController
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
go test ./internal/middleware -v
```

## Testing

The project uses Go's built-in testing framework with:
- **Unit Tests:** Each package has corresponding `*_test.go` files
- **Mocks:** Mock implementations of interfaces in `mocks.go` files for isolated testing
- **Middleware Tests:** Comprehensive tests for auth middleware in `internal/middleware/auth_test.go`

Run all tests with coverage:
```bash
go test ./... -cover
```
