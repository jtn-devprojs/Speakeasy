import 'package:get_it/get_it.dart';
import 'services.dart';

/// Dependency Injection Container setup
final getIt = GetIt.instance;

/// Configure all dependencies for the application
void setupDependencies() {
  // Register services
  getIt.registerSingleton<ApiService>(ApiService());
  getIt.registerSingleton<UserService>(UserService(getIt<ApiService>()));
  getIt.registerSingleton<AuthService>(AuthService(getIt<UserService>()));
}

/// Retrieve a service from the DI container
T get<T extends Object>() => getIt<T>();
