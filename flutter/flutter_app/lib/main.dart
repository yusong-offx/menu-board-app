import 'package:flutter/material.dart';
import 'page/home.dart';

const myblack = Color(0xff181818);
const mywhite = Color(0xffFAFAFA);
const myblue = Color(0xff81ECE6);
const myorange = Color(0xffF2B33A);

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Material App!",
      theme: ThemeData(
        backgroundColor: Colors.black,
      ),
      home: Home(),
      routes: {
        "/": (context) => Home(),
      },
    );
  }
}
