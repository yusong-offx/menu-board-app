import 'package:flutter/material.dart';

class Base extends StatelessWidget {
  const Base({
    Key? key,
  }) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Text("main page ${MediaQuery.of(context).size}"),
    );
  }
}
