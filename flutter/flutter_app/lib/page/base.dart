import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:flutter_app/main.dart';

class Base extends ConsumerStatefulWidget {
  const Base({super.key});

  @override
  ConsumerState<Base> createState() => _BaseState();
}

class _BaseState extends ConsumerState<Base> {
  // @override
  // void initState() {
  //   super.initState();
  //   ref.read(counterProiver);
  // }

  @override
  Widget build(BuildContext context) {
    final counter = ref.watch(myStateProvider);
    return const Text("good");
    // return ListView.separated(
    //   itemCount: itemCount,
    //   separatorBuilder: separatorBuilder,
    //   itemBuilder: itemBuilder,
    // );
  }
}
