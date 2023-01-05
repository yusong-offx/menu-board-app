import 'package:flutter/material.dart';

class LoginTextField extends StatelessWidget {
  final String lable;
  final ctl = TextEditingController();
  bool star = false;

  LoginTextField({
    Key? key,
    required this.lable,
  }) : super(key: key);

  LoginTextField.forPassword({Key? key, required this.lable}) : star = true;

  @override
  Widget build(BuildContext context) {
    return TextField(
      obscureText: star,
      controller: ctl,
      decoration: InputDecoration(
        labelText: lable,
      ),
    );
  }
}
