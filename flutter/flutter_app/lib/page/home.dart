import 'package:flutter/material.dart';

class Home extends StatelessWidget {
  const Home({super.key});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: const Text("leading"),
        title: Column(children: const [
          Text(
            "Hey, Eugene",
            style: TextStyle(fontSize: 28, fontWeight: FontWeight.w700),
          ),
          Text(
            "Welecome back",
            style: TextStyle(fontSize: 18),
          )
        ]),
      ),
      body: Center(
        key: UniqueKey(),
        child: const Text("This is home body"),
      ),
    );
  }
}
