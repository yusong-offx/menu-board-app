import 'package:flutter/material.dart';
import 'package:flutter_app/widget/logintextfield.dart';
import 'package:flutter_app/widget/signupcontainer.dart';
import 'package:flutter_app/main.dart';

class SignUp extends StatefulWidget {
  const SignUp({Key? key}) : super(key: key);

  @override
  State<SignUp> createState() => _SignUpState();
}

class _SignUpState extends State<SignUp> {
  final LoginTextField id = LoginTextField(lable: "ID");
  final LoginTextField password =
      LoginTextField.fromPassword(lable: "Password");
  final LoginTextField firstName = LoginTextField(lable: "First Name");
  final LoginTextField lastName = LoginTextField(lable: "Last Name");
  final LoginTextField email = LoginTextField(lable: "Email");

  final snackbar = const SnackBar(
    duration: Duration(milliseconds: 600),
    content: Text(
      "Success to Join us!",
      style: TextStyle(
        color: Colors.white,
        fontSize: 18,
      ),
    ),
    behavior: SnackBarBehavior.floating,
  );
  bool idValidation = false;
  bool postOK = false;
  bool buttonActive = true;

  void onJoin() async {
    // Form Validation Check
    setState(() {
      buttonActive = false;
    });
    postOK = await api.postUser({
      "login_id": id.ctl.text,
      "login_password": password.ctl.text,
      "first_name": firstName.ctl.text,
      "last_name": lastName.ctl.text,
      "email": email.ctl.text,
    });
    setState(() {
      if (postOK) {
        Navigator.pop(context);
        ScaffoldMessenger.of(context).showSnackBar(snackbar);
      }
      buttonActive = true;
    });
  }

  void isIDValid() {
    // check invalid
    setState(() {
      Navigator.pop(context);
      ScaffoldMessenger.of(context).showSnackBar(snackbar);
      idValidation = !idValidation;
    });
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
                child: buttonActive
                    ? TextButton(
                        onPressed: onJoin,
                        child: Text(
                          "Join",
                          style: Theme.of(context).textTheme.titleLarge,
                        ),
                      )
                    : const CircularProgressIndicator(),
              )
            ],
          ),
        ),
      ),
    );
  }
}
