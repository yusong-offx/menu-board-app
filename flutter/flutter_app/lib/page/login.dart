import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:flutter_app/widget/logintextfield.dart';

class Login extends StatelessWidget {
  final LoginTextField id = LoginTextField(lable: "ID");
  final LoginTextField password =
      LoginTextField.fromPassword(lable: "Password");

  Login({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final deviceSize = MediaQuery.of(context).size;
    return Scaffold(
      body: SingleChildScrollView(
        padding: EdgeInsets.fromLTRB(deviceSize.width * 0.1,
            deviceSize.height * 0.15, deviceSize.width * 0.1, 0),
        child: Column(children: [
          Center(
            child: Text(
              "🍖 Enjoy Your Meal 🍖",
              style: TextStyle(
                fontSize: 28,
                color: const Color(0xFF442c42).withOpacity(0.8),
              ),
            ),
          ),
          const SizedBox(height: 30),
          Container(
            padding: const EdgeInsets.symmetric(
              vertical: 10,
              horizontal: 65,
            ),
            decoration: BoxDecoration(
              borderRadius: const BorderRadius.all(
                Radius.circular(10),
              ),
              border: Border.all(
                width: 1,
                color: Colors.black26,
              ),
            ),
            child: Column(children: [
              SvgPicture.asset(
                "assets/images/loginlogo.svg",
                height: deviceSize.height * 0.25,
                color: const Color(0xFF442c42).withOpacity(0.8),
              ),
              id,
              const SizedBox(height: 10),
              password,
              const SizedBox(height: 20),
              TextButton(
                style: ButtonStyle(
                  minimumSize: MaterialStateProperty.all(
                    const Size(200, 42),
                  ),
                ),
                onPressed: () {
                  // Navigator.pushNamedAndRemoveUntil(
                  //     context, "/base", (route) => false);
                  Navigator.pushNamed(context, "/base");
                },
                child: Text(
                  "Login",
                  style: Theme.of(context).textTheme.titleLarge,
                ),
              ),
            ]),
          ),
          const SizedBox(height: 30),
        ]),
      ),
    );
  }
}
