import 'package:flutter/material.dart';
import 'package:flutter_app/model/apiservice.dart';
import 'package:flutter_app/page/home.dart';
import 'package:flutter_app/page/qrcode.dart';
import 'package:flutter_app/page/signup.dart';
import 'package:flutter_app/page/login.dart';
import 'package:flutter_app/page/base.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

final api = APIService();

class MyCounter extends StateNotifier<String> {
  MyCounter() : super("hey world");

  void set(String val) {
    super.state = val;
  }
}

var counterProiver = StateNotifierProvider<MyCounter, String>(
  (ref) => MyCounter(),
);

void main() {
  runApp(const ProviderScope(
    child: MyApp(),
  ));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        backgroundColor: const Color(0xFFb07845),
        textTheme: TextTheme(
          titleLarge: TextStyle(
            fontSize: 22,
            color: const Color(0xFF442c42).withOpacity(0.8),
          ),
          titleMedium: const TextStyle(
            fontSize: 18,
            color: Color(0xFF21233a),
          ),
          bodyLarge: const TextStyle(color: Colors.black87),
        ),
      ),
      home: Container(
        color: const Color(0xFFb07845),
        child: const Home(),
      ),
      routes: {
        "/home": (context) => const Home(),
        "/login": (context) => Login(),
        "/signup": (context) => const SignUp(),
        "/base": (context) => const Base(),
        "/qrcode": (context) => const QRcode(),
      },
    );
  }
}



// const myblack = Color(0xFFb07845);
// const mywhite = Color(0xFF21233a);
// const myblue = Color(0xFF5b1829);wwwwwwwwwwwwwwwwwwwwww
// const asd = Color(0xFF442c42);
// const mygreen = Color(0xFF0d5240);