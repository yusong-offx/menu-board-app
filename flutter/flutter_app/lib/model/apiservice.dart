import 'dart:convert';

import 'package:http/http.dart' as http;

class APIService {
  static const baseUrl = "http://yusong-offx.link";

  Future<bool> postUser(Map<String, dynamic> data) async {
    var resp = await http.post(
      Uri.parse("$baseUrl/user/login"),
      body: jsonEncode(data),
    );
    if (resp.statusCode == 201) {
      return true;
    }
    return false;
  }
}
