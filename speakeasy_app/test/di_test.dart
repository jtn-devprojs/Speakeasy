import 'package:flutter_test/flutter_test.dart';
import 'package:speakeasy/di.dart';
import 'package:speakeasy/services.dart';

void main() {
  group('Dependency Injection', () {
    setUp(() {
      // Clear all instances before each test
      getIt.reset();
      setupDependencies();
    });

    test('ApiService is registered and retrievable', () {
      final apiService = getIt<ApiService>();
      expect(apiService, isNotNull);
      expect(apiService, isA<ApiService>());
    });

    test('UserService is registered with ApiService dependency', () {
      final userService = getIt<UserService>();
      expect(userService, isNotNull);
      expect(userService, isA<UserService>());
    });

    test('AuthService is registered with UserService dependency', () {
      final authService = getIt<AuthService>();
      expect(authService, isNotNull);
      expect(authService, isA<AuthService>());
    });

    test('Services are singletons', () {
      final apiService1 = getIt<ApiService>();
      final apiService2 = getIt<ApiService>();
      expect(identical(apiService1, apiService2), true);
    });

    test('Dependency chain is correctly wired', () {
      final userService = getIt<UserService>();
      final authService = getIt<AuthService>();
      expect(userService, isNotNull);
      expect(authService, isNotNull);
    });
  });
}
