# Speakeasy App

A cross-platform mobile and web application built with Flutter for Android, iOS, and web platforms.

## Architecture

This project follows the **Dependency Injection** pattern using the `get_it` package for service management and loose coupling.

### Folder Structure

```
speakeasy_app/
├── lib/
│   ├── main.dart          # Application entry point
│   ├── app.dart           # Root widget and app configuration
│   ├── di.dart            # Dependency injection setup
│   ├── services.dart      # Business logic services
│   └── ...
├── test/
│   ├── di_test.dart       # DI container tests
│   ├── services_test.dart # Service tests
│   └── ...
├── pubspec.yaml           # Dependencies and metadata
└── README.md              # This file
```

## Services

### ApiService
Handles all HTTP communication with the Speakeasy backend API.
- `get()` - Make GET requests
- `post()` - Make POST requests
- `put()` - Make PUT requests
- `delete()` - Make DELETE requests

### UserService
Manages user-related operations.
- `getUserProfile()` - Fetch user profile
- `updateUserProfile()` - Update user profile
- `getUserPreferences()` - Get user preferences

### AuthService
Handles authentication and token management.
- `login()` - Authenticate user
- `logout()` - End user session
- `register()` - Create new user account
- `refreshToken()` - Refresh authentication token
- `isAuthenticated()` - Check authentication status

## Dependency Injection Setup

Services are registered as singletons in `lib/di.dart`:

```dart
void setupDependencies() {
  getIt.registerSingleton<ApiService>(ApiService());
  getIt.registerSingleton<UserService>(UserService(getIt<ApiService>()));
  getIt.registerSingleton<AuthService>(AuthService(getIt<UserService>()));
}
```

Access services anywhere in the app:
```dart
final authService = getIt<AuthService>();
```

## Getting Started

### Prerequisites
- Flutter SDK 3.0.0 or higher
- Dart 3.0.0 or higher

### Installation

1. Navigate to the project directory:
```bash
cd speakeasy_app
```

2. Install dependencies:
```bash
flutter pub get
```

3. Run the app:
```bash
flutter run
```

### Running Tests

Run all tests:
```bash
flutter test
```

Run specific test file:
```bash
flutter test test/di_test.dart
```

## Implementation Status

All service methods are currently stubbed with `TODO` markers. Implementation will follow in subsequent phases.

- [ ] ApiService HTTP methods
- [ ] UserService user operations
- [ ] AuthService authentication flows
- [ ] UI implementation
- [ ] State management integration
