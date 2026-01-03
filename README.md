# Speakeasy - Multi-Platform App

A modern, cross-platform application built with Flutter (frontend) and Go (backend), designed for Android, iOS, and web platforms.

## Project Structure

```
Speakeasy/
├── speakeasy_app/         # Flutter app (Dart)
│   ├── lib/
│   │   ├── main.dart      # Application entry point
│   │   ├── app.dart       # Root widget and app configuration
│   │   ├── di.dart        # Dependency injection setup
│   │   ├── services.dart  # Business logic services
│   │   └── ...
│   ├── test/              # Unit tests
│   │   ├── di_test.dart
│   │   ├── services_test.dart
│   │   └── ...
│   ├── pubspec.yaml       # Dependencies
│   └── README.md          # App documentation
│
├── speakeasy_api/         # Go API server with Gin framework
│   ├── cmd/
│   │   └── server/
│   │       └── main.go    # Server entry point
│   ├── internal/          # Internal packages
│   │   ├── di/
│   │   │   ├── container.go
│   │   │   └── di_test.go
│   │   ├── controllers/
│   │   │   ├── user_controller.go
│   │   │   ├── auth_controller.go
│   │   │   └── session_controller.go
│   │   └── services/
│   │       ├── user_service.go
│   │       ├── user_service_test.go
│   │       ├── auth_service.go
│   │       ├── auth_service_test.go
│   │       ├── session_service.go
│   │       └── errors.go
│   ├── go.mod             # Go module definition
│   └── README.md          # API documentation
│
└── README.md              # This file
```

## Quick Start

### API (Go)

```bash
cd speakeasy_api
go mod download
go run cmd/server/main.go
```

Server runs on `http://localhost:8080`

### App (Flutter)

```bash
cd speakeasy_app
flutter pub get
flutter run
```

### Run Tests

**API:**
```bash
cd speakeasy_api
go test ./...
```

**App:**
```bash
cd speakeasy_app
flutter test
```

## Architecture Overview

Both projects use the **Dependency Injection** pattern for:
- **Loose coupling** between components
- **Easy testing** through mock injection
- **Centralized configuration** management
- **Clear dependency flow**

### API Architecture

**Data Layer**: Repositories handle all database operations with clean interfaces
- `IUserRepository` - User CRUD operations
- `ISessionRepository` - Session management
- `ISessionUserRepository` - Session membership tracking
- `IMessageRepository` - Message operations

**Business Layer**: Services contain business logic and transactions
- `IUserService` - User business logic
- `IAuthService` - Authentication logic
- `ISessionService` - Session operations with location handling

**HTTP Layer**: Controllers expose endpoints via Gin framework
- `UserController` - User endpoints
- `AuthController` - Auth endpoints
- `SessionController` - Session endpoints

**DI Container**: Located in `speakeasy_api/internal/di/container.go`, manages all dependencies

### App DI Container
Located in `speakeasy_app/lib/di.dart`, uses `get_it` package for service registration.

## Technology Stack

### App
- **Framework**: Flutter
- **Language**: Dart 3.0+
- **DI Container**: get_it
- **Testing**: flutter_test, mockito

### API
- **Language**: Go 1.24+
- **Framework**: Gin (HTTP web framework)
- **Database**: SQLite (in-memory for development/testing)
- **DI Pattern**: Manual constructor injection with interface-based contracts
- **Testing**: Go testing package with mocks

### Database Schema (Go Backend)
- **sessions**: Stores location-based chat sessions with coordinates, accuracy, and timestamps
- **session_users**: Tracks user participation in sessions with join/leave times
- **users**: User accounts and profiles
- **messages**: Chat messages within sessions

## Development Workflow

1. **Implement service methods** in respective service files
2. **Implement controllers** that use the services
3. **Write unit tests** as you implement features
4. **Register new services** in the DI container
5. **Connect frontend** to backend endpoints

## Implementation Status

### API Services & Data Access
- [x] User repository and service
- [x] Session repository and service  
- [x] SessionUser repository (membership tracking)
- [x] Message repository
- [x] Dependency injection pattern
- [x] Database schema design
- [ ] Location-based proximity matching (Haversine formula)
- [ ] Session lifecycle management
- [ ] Authentication (JWT tokens)
- [ ] Request validation
- [ ] Error handling middleware

### App Services
- [ ] API client HTTP methods
- [ ] Location services (GPS integration)
- [ ] Authentication flow
- [ ] User profile management
- [ ] Session discovery and joining
- [ ] Chat messaging UI
- [ ] State management
- [ ] Location permissions handling

## API Documentation

See [speakeasy_api/README.md](speakeasy_api/README.md) for detailed API endpoint documentation.

## Code Examples

### Using Services in Flutter

```dart
// Get service from container
final authService = getIt<AuthService>();

// Use the service
String token = await authService.login('user', 'password');
```

### Using Services in Go

```go
// Create DI container with all dependencies
container := di.NewContainer(db)

// Access services
user, err := container.UserService.GetUserByID("123")
sessions, err := container.SessionService.GetNearbyLocations(lat, lon, radiusKm)

// Access repositories for direct data access
users, err := container.SessionUserRepo.GetActiveUsersInSession("session-id")
```

## Contributing

1. Implement stubs marked with `TODO`
2. Write tests for new functionality
3. Follow existing code patterns
4. Update README with new features

## Next Steps

### Backend (Go API)
1. Implement location-based proximity calculations (Haversine formula)
2. Add JWT token authentication
3. Implement session lifecycle (create, join, leave, end)
4. Add geolocation API integration (optional: reverse geocoding)
5. Implement message persistence and retrieval
6. Add request validation and error handling middleware
7. Set up environment configuration (port, database, etc.)
8. Add API documentation (Swagger/OpenAPI)

### Frontend (Flutter App)
1. Implement GPS location services
2. Add authentication UI and flow
3. Create session discovery UI
4. Build chat interface
5. Implement real-time message updates
6. Add location permissions handling
7. Connect to backend API endpoints
8. Implement offline-first caching