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
│   │   ├── config_loader.go             # Configuration management (loads from .env files)
│   │   ├── config_loader_test.go        # Config loader unit tests
│   │   ├── interfaces.go                # Config interfaces
│   │   └── .env.dev                     # Test configuration file
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
├── .env.dev                             # Development environment settings
├── .env.prod                            # Production environment settings
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
- Firebase CLI (optional, for emulator management)

### Configuration

Configuration is managed via environment-specific settings files:

- **`.env.dev`** - Development settings (default fallback)
  - `PORT=8080` - Server port
  - `DB_TYPE=sqlite` - Database type
  - `DB_CONNECTION=:memory:` - In-memory SQLite for development
  - `FIREBASE_EMULATOR_HOST=localhost:9099` - Firebase emulator endpoint

- **`.env.prod`** - Production settings
  - Database connection to production MySQL server
  - Firebase authentication to production

### Installation

1. Navigate to the project directory:
```powershell
cd speakeasy_api
```

2. Download dependencies:
```powershell
go mod download
```

3. Run the server (uses `.env.dev` by default):
```powershell
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### Firebase Emulator Setup

For local development, Firebase Authentication uses the emulator instead of production:

1. **Install Firebase CLI** (if not already installed):

   Option A - Download Windows installer (recommended):
   - Visit https://firebase.google.com/docs/cli#install-cli-windows and download the installer
   
   Option B - Use npm (requires Node.js):
```powershell
npm install -g firebase-tools
```

2. **Start Firebase Emulator Suite** (in the project root):
```powershell
firebase emulators:start
```

The Firebase Authentication emulator will start on `localhost:9099` (configured in `.env.dev`)

3. **Configure your Flutter app** to use the emulator:
   - Set the Firebase emulator host to `localhost:9099` in your app configuration
   - See `speakeasy_app/README.md` for Flutter-specific emulator setup

### Running with Custom Environment

To run with a different environment setting:

```powershell
# Production environment (if .env.prod exists)
$env:ENVIRONMENT = "prod"
go run cmd/server/main.go
```

To override specific settings:

```powershell
# Override via environment variables
$env:PORT = "9000"
$env:DB_TYPE = "postgres"
$env:DB_CONNECTION = "user:pass@localhost:5432/speakeasy"
go run cmd/server/main.go
```

### Running Tests

Run all tests:
```powershell
go test ./...
```

Run tests with verbose output:
```powershell
go test -v ./...
```

Run tests with coverage report:
```powershell
go test ./... -cover
```

Run specific test package:
```powershell
go test ./internal/middleware -v
```

Tests use the shared `.env.dev` settings file from `internal/config/` directory.

## Configuration Details

### Settings File Hierarchy

Configuration is loaded strictly from the environment-specific file:

- `.env.{environment}` (e.g., `.env.dev`, `.env.prod`)

**Important:** Settings files are the source of truth. All configuration values (PORT, DB_TYPE, DB_CONNECTION) must be defined in the settings file. The only environment variable used is `ENVIRONMENT` to determine which settings file to load.

### Required Settings

All settings must be defined in the `.env.{environment}` file. No hardcoded defaults exist, and no environment variables override file settings.

| Setting | Example | Required | Source |
|---------|---------|----------|--------|
| `PORT` | `8080` | Yes | Settings file |
| `DB_TYPE` | `sqlite` or `mysql` | Yes | Settings file |
| `DB_CONNECTION` | `:memory:` or connection string | Yes | Settings file |
| `FIREBASE_EMULATOR_HOST` | `localhost:9099` | Only for dev with Firebase emulator | Settings file |
| `ENVIRONMENT` | `dev` or `prod` | No (defaults to `dev`) | Environment variable or default |

### Configuration Example

**Development (.env.dev):**
```
PORT=8080
DB_TYPE=sqlite
DB_CONNECTION=:memory:
FIREBASE_EMULATOR_HOST=localhost:9099
```

**Production (.env.prod):**
```
PORT=8080
DB_TYPE=mysql
DB_CONNECTION=user:password@tcp(prod-db:3306)/speakeasy
```

### Switching Environments

**Development (default):**
```powershell
# Loads .env.dev automatically
go run cmd/server/main.go
```

**Production:**
```powershell
# Set ENVIRONMENT to "prod" to load .env.prod
$env:ENVIRONMENT = "prod"
go run cmd/server/main.go
```

**Key Points:**
- The `ENVIRONMENT` variable determines which `.env.{environment}` file to load
- All configuration values come from the settings file, not environment variables
- The required `.env.{environment}` file must exist - there is no fallback

## Testing

The project uses Go's built-in testing framework with:
- **Unit Tests:** Each package has corresponding `*_test.go` files
- **Mocks:** Mock implementations of interfaces in `mocks.go` files for isolated testing
- **Middleware Tests:** Comprehensive tests for auth middleware in `internal/middleware/auth_test.go`
- **Config Tests:** Configuration loading tests in `internal/config/config_loader_test.go`
- **Test Settings:** Uses `.env.dev` in `internal/config/` directory for test configuration

Run all tests with coverage:
```bash
go test ./... -cover
```

## Future Enhancements

- **Configuration File Validator:** Validate configuration files for invalid DB_TYPE values, invalid port ranges, malformed connection strings, and other common configuration errors at startup
- **Secret Management Integration:** Move database credentials from configuration files to a secret management service (e.g., AWS Secrets Manager, HashiCorp Vault) for improved security
