import 'package:flutter_test/flutter_test.dart';
import 'package:speakeasy/services.dart';

void main() {
  group('ApiService', () {
    late ApiService apiService;

    setUp(() {
      apiService = ApiService();
    });

    test('ApiService initializes correctly', () {
      expect(apiService, isNotNull);
    });

    test('get method is defined', () async {
      expect(
        () => apiService.get('/test'),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('post method is defined', () async {
      expect(
        () => apiService.post('/test', {}),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('put method is defined', () async {
      expect(
        () => apiService.put('/test', {}),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('delete method is defined', () async {
      expect(
        () => apiService.delete('/test'),
        throwsA(isA<UnimplementedError>()),
      );
    });
  });

  group('UserService', () {
    late UserService userService;
    late ApiService apiService;

    setUp(() {
      apiService = ApiService();
      userService = UserService(apiService);
    });

    test('UserService initializes with ApiService', () {
      expect(userService, isNotNull);
    });

    test('getUserProfile method is defined', () async {
      expect(
        () => userService.getUserProfile('user123'),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('updateUserProfile method is defined', () async {
      expect(
        () => userService.updateUserProfile('user123', {}),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('getUserPreferences method is defined', () async {
      expect(
        () => userService.getUserPreferences('user123'),
        throwsA(isA<UnimplementedError>()),
      );
    });
  });

  group('AuthService', () {
    late AuthService authService;
    late UserService userService;
    late ApiService apiService;

    setUp(() {
      apiService = ApiService();
      userService = UserService(apiService);
      authService = AuthService(userService);
    });

    test('AuthService initializes with UserService', () {
      expect(authService, isNotNull);
    });

    test('login method is defined', () async {
      expect(
        () => authService.login('user', 'pass'),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('logout method is defined', () async {
      expect(
        () => authService.logout(),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('register method is defined', () async {
      expect(
        () => authService.register('user', 'email', 'pass'),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('refreshToken method is defined', () async {
      expect(
        () => authService.refreshToken('token'),
        throwsA(isA<UnimplementedError>()),
      );
    });

    test('isAuthenticated method is defined', () async {
      expect(
        () => authService.isAuthenticated(),
        throwsA(isA<UnimplementedError>()),
      );
    });
  });
}
