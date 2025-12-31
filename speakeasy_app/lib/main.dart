import 'package:flutter/material.dart';
import 'di.dart';
import 'app.dart';

void main() {
  // Initialize dependency injection
  setupDependencies();

  runApp(const SpeakeasyApp());
}
