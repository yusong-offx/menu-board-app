import 'package:flutter/material.dart';
import 'package:flutter_app/widget/logintextfield.dart';
import 'package:flutter_app/main.dart';
import 'package:flutter_app/widget/signupcontainer.dart';
import 'package:top_snackbar_flutter/custom_snack_bar.dart';
import 'package:top_snackbar_flutter/top_snack_bar.dart';

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

  bool idValidation = false;
  bool buttonActive = true;

  void topSnackbar(String msg, bool err) {
    showTopSnackBar(
      Overlay.of(context)!,
      err
          ? CustomSnackBar.error(
              message: msg,
              textStyle: const TextStyle(color: Colors.white, fontSize: 22),
            )
          : CustomSnackBar.success(
              message: msg,
              textStyle: const TextStyle(color: Colors.white, fontSize: 22),
            ),
      displayDuration: const Duration(milliseconds: 100),
    );
  }

  void onJoin() async {
    if (!idValidation) {
      topSnackbar(
        "Check your ID",
        true,
      );
      return;
    } else if (password.ctl.text.isEmpty) {
      topSnackbar(
        "Write your Password",
        true,
      );
      return;
    } else if (firstName.ctl.text.isEmpty) {
      topSnackbar(
        "Write your First Name",
        true,
      );
      return;
    } else if (lastName.ctl.text.isEmpty) {
      topSnackbar(
        "Write your Last Name",
        true,
      );
      return;
    } else if (email.ctl.text.isEmpty) {
      topSnackbar(
        "Write your Email",
        true,
      );
      return;
    }
    setState(() {
      buttonActive = false;
    });
    bool ok = await api.signUpUser({
      "login_id": id.ctl.text,
      "login_password": password.ctl.text,
      "first_name": firstName.ctl.text,
      "last_name": lastName.ctl.text,
      "email": email.ctl.text,
    });
    setState(
      () {
        if (ok) {
          Navigator.pop(context);
          topSnackbar(
            "Success to Join us!",
            false,
          );
          return;
        }
        buttonActive = true;
        topSnackbar(
          "Something Wrong",
          true,
        );
      },
    );
  }

  void isIDValid() async {
    String inputID = id.ctl.text;
    if (inputID.isEmpty) {
      topSnackbar("Write your ID", true);
      return;
    }
    idValidation = await api.signUpIDChecker(inputID);
    if (!idValidation) {
      topSnackbar("User exist", true);
    }
    setState(() {});
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
