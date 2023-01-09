import 'package:flutter/material.dart';
import 'package:flutter_app/main.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

class QRcode extends ConsumerWidget {
  const QRcode({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return const Center(
      child: Text("qrcode"),
    );
  }
}
