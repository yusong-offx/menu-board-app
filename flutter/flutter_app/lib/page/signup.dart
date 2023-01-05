import 'package:flutter/material.dart';
import 'package:flutter_app/widget/logintextfield.dart';
import 'package:flutter_app/widget/signupcontainer.dart';

class SignUp extends StatefulWidget {
  const SignUp({Key? key}) : super(key: key);

  @override
  State<SignUp> createState() => _SignUpState();
}

class _SignUpState extends State<SignUp> {
  final LoginTextField id = LoginTextField(lable: "ID");
  final LoginTextField password = LoginTextField(lable: "Password");
  final LoginTextField firstName = LoginTextField(lable: "First Name");
  final LoginTextField lastName = LoginTextField(lable: "Last Name");
  final LoginTextField email = LoginTextField(lable: "Email");

  bool idValidation = false;

  void isIDValid() {
    setState(() {
      idValidation = !idValidation;
    });
    // check invalid
  }

  @override
  Widget build(BuildContext context) {
    final deviceSize = MediaQuery.of(context).size;
    return Scaffold(
      body: SizedBox(
        height: deviceSize.height,
        width: deviceSize.width,
        child: SingleChildScrollView(
          padding: EdgeInsets.symmetric(
            vertical: deviceSize.height * 0.1,
            horizontal: deviceSize.width * 0.1,
          ),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(
                height: 100,
                child: Center(
                  child: Text(
                    "\u{1F32D} Welcome \u{1F32D}",
                    style: TextStyle(
                      fontSize: 30,
                    ),
                  ),
                ),
              ),
              SignUpContainer(children: [
                Text(
                  "Login Info",
                  style: Theme.of(context).textTheme.titleLarge,
                ),
                Row(
                  crossAxisAlignment: CrossAxisAlignment.end,
                  children: [
                    Flexible(
                      flex: 10,
                      child: id,
                    ),
                    Flexible(
                      flex: 2,
                      child: IconButton(
                        icon: Icon(
                          idValidation
                              ? Icons.sentiment_very_satisfied
                              : Icons.sentiment_satisfied,
                          color: const Color(0xFF442c42).withOpacity(0.8),
                          size: 35,
                        ),
                        onPressed: isIDValid,
                      ),
                    )
                  ],
                ),
                password,
              ]),
              const SizedBox(height: 20),
              SignUpContainer(
                children: [
                  Text(
                    "Personal Info",
                    style: Theme.of(context).textTheme.titleLarge,
                  ),
                  firstName,
                  lastName,
                  email,
                ],
              ),
              const SizedBox(height: 20),
              Center(
                child: TextButton(
                  onPressed: () {},
                  child: Text(
                    "Join",
                    style: Theme.of(context).textTheme.titleLarge,
                  ),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
