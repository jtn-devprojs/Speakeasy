# Speakeasy App

A cross-platform mobile and web application built with Flutter for Android, iOS, and web platforms. Handles user authentication via Firebase and communicates with the Speakeasy backend API for session and location management.

## Architecture

This project follows the **Dependency Injection** pattern using the `get_it` package for service management and loose coupling.

### Authentication Flow

1. **Client-Side Authentication (Firebase):** 
   - User signs in/up with Firebase directly in the app
   - Firebase returns an ID token
   - Token is stored locally and included in all API requests

2. **Server-Side Token Validation:**
   - Backend validates the Firebase token on each request
   - No user credentials are transmitted to the backend
   - Server handles sessions, locations, and preferences only

### Folder Structure

```
speakeasy_app/
├── lib/
│   ├── main.dart          # Application entry point
│   ├── app.dart           # Root widget and app configuration
│   ├── di.dart            # Dependency injection setup (get_it)
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

### FirebaseAuthService
Handles authentication with Firebase.
- `signUp()` - Create new user account
- `signIn()` - Authenticate user
- `signOut()` - Logout user
- `getToken()` - Get Firebase ID token
- `isAuthenticated()` - Check if user is logged in

### ApiService
Handles all HTTP communication with the Speakeasy backend API.
- `get()` - Make GET requests (includes Firebase token)
- `post()` - Make POST requests (includes Firebase token)
- `put()` - Make PUT requests (includes Firebase token)
- `delete()` - Make DELETE requests (includes Firebase token)

**Note:** All requests automatically include the Firebase ID token in the Authorization header.

### SessionService
Manages session-related operations.
- `getNearbySessions()` - Get sessions near user location
- `joinSession()` - Join a session
- `leaveSession()` - Leave a session

### LocationService
Handles location tracking.
- `getUserLocation()` - Get current user location
- `updateUserLocation()` - Update user location
- `checkVicinity()` - Check if user is near a location

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
