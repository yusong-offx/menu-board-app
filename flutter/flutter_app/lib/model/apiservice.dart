import 'dart:convert';

import 'package:http/http.dart' as http;

class APIService {
  static const baseUrl = "http://yusong-offx.link";
  Map<String, String> headers = {"Content-Type": "application/json"};

  Future<bool> signUpUser(Map<String, dynamic> data) async {
    var resp = await http.post(
      Uri.parse("$baseUrl/user/signup"),
      headers: headers,
      body: jsonEncode(data),
    );
    if (resp.statusCode == 201) {
      return true;
    }
    return false;
  }

  Future<bool> signUpIDChecker(String id) async {
    var resp = await http.get(Uri.parse("$baseUrl/user/signup/$id"));
    if (resp.statusCode == 200) {
      return true;
    }
    return false;
  }
}
