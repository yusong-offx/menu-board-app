import 'package:flutter/material.dart';
import 'package:flutter_app/page/main.dart';

class Home extends StatelessWidget {
  const Home({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(50),
      child: Column(
        children: [
          Text(
            "큰 제목입니다.",
            // style: TextStyle(fontSize: 50),
            style: Theme.of(context).textTheme.titleLarge,
          ),
          Text(
            "작은 제목입니다.",
            style: Theme.of(context).textTheme.titleMedium,
          ),
          const Main()
        ],
      ),
    );
  }
}
