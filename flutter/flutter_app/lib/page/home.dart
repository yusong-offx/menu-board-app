import 'package:flutter/material.dart';

class Home extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final deviceSize = MediaQuery.of(context).size;
    return Scaffold(
      body: Padding(
        padding: EdgeInsets.symmetric(
          vertical: deviceSize.width * 0.2,
          horizontal: deviceSize.height * 0.05,
        ),
        child: Column(
          children: [
            Flexible(
              flex: 8,
              child: Container(
                width: deviceSize.width,
                height: deviceSize.height,
                decoration: const BoxDecoration(
                  color: Colors.black,
                ),
                child: const Text("camera"),
              ),
            ),
            const SizedBox(height: 30),
            Flexible(
              flex: 1,
              child: OutlinedButton(
                style: ButtonStyle(
                  minimumSize: MaterialStateProperty.all(
                    const Size(200, 42),
                  ),
                ),
                onPressed: () {
                  Navigator.pushNamed(context, "login");
                  //   Navigator.pushNamedAndRemoveUntil(
                  //       context, "login", (route) => false);
                },
                child: Text(
                  "Login",
                  style: Theme.of(context).textTheme.titleLarge,
                ),
              ),
            ),
            const SizedBox(height: 20),
            Flexible(
              flex: 1,
              child: OutlinedButton(
                style: ButtonStyle(
                  minimumSize: MaterialStateProperty.all(
                    const Size(200, 42),
                  ),
                ),
                onPressed: () {
                  Navigator.pushNamed(context, "signup");
                },
                child: Text(
                  "Sign Up",
                  style: Theme.of(context).textTheme.titleLarge,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
