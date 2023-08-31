import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:warrant/counter/view/warrant_page.dart';
import 'package:warrant/warrant_bloc.dart';

void main() => runApp(const WarrantApp());

class WarrantApp extends StatelessWidget {
  const WarrantApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: BlocProvider(
        create: (_) => WarrantBloc(),
        child: const WarrantPage(),
      ),
    );
  }
}
