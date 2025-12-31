# Speakeasy - Multi-Platform App

A modern, cross-platform application built with Flutter (frontend) and Go (backend), designed for Android, iOS, and web platforms.

## Project Structure

```
Speakeasy/
├── speakeasy_app/         # Flutter app (Dart)
│   ├── lib/               # Application source code
│   ├── test/              # Unit tests
│   └── pubspec.yaml       # Dependencies
│
├── speakeasy_api/         # Go API server
│   ├── cmd/server/        # Server entry point
│   ├── internal/          # Internal packages
│   │   ├── di/            # Dependency injection
│   │   ├── handlers/      # HTTP handlers
│   │   └── services/      # Business logic
│   ├── test/              # Unit tests
│   └── go.mod             # Go module definition
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

### API DI Container
Located in `speakeasy_api/internal/di/container.go`, manages all services and handlers.

### App DI Container
Located in `speakeasy_app/lib/di.dart`, uses `get_it` package for service registration.

## Technology Stack

### App
- **Framework**: Flutter
- **Language**: Dart 3.0+
- **DI Container**: get_it
- **Testing**: flutter_test, mockito

### API
- **Language**: Go 1.21+
- **Router**: Gorilla Mux
- **DI Pattern**: Manual container
- **Testing**: Go testing package

## Development Workflow

1. **Implement service methods** in respective service files
2. **Implement handlers** that use the services
3. **Write unit tests** as you implement features
4. **Register new services** in the DI container
5. **Connect frontend** to backend endpoints

## Implementation Status

### API Services
- [ ] User management (create, read, update, delete)
- [ ] Authentication (login, logout, token management)
- [ ] User preferences
- [ ] Database integration
- [ ] Request validation
- [ ] Error handling middleware

### App Services
- [ ] API client HTTP methods
- [ ] Authentication flow
- [ ] User profile management
- [ ] Preference synchronization
- [ ] State management
- [ ] UI implementation

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
// Create DI container
container := di.NewContainer()

// Access services
user, err := container.UserService.GetUserByID("123")
```

## Contributing

1. Implement stubs marked with `TODO`
2. Write tests for new functionality
3. Follow existing code patterns
4. Update README with new features

## Next Steps

1. Implement database models and repositories
2. Add JWT token authentication
3. Set up environment configuration
4. Add API request/response validation
5. Implement error handling middleware
6. Connect frontend to backend
7. Add comprehensive test coverage