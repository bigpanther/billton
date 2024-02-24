// import 'package:flutter/material.dart';
// import 'package:flutter_bloc/flutter_bloc.dart';
// import 'package:warrant/warrant/view/warrant_page.dart';
// import 'package:warrant/warrant_bloc.dart';
// import 'package:firebase_auth/firebase_auth.dart';

// void main() => runApp(const WarrantApp());

// class WarrantApp extends StatelessWidget {
//   const WarrantApp({super.key});

//   @override
//   Widget build(BuildContext context) {
//     return MaterialApp(
//       home: BlocProvider(
//         create: (_) => WarrantBloc(),
//         child: const GoogleSignInScreen(),
//       ),
//     );
//   }
// }
import 'package:flutter/material.dart';
import 'package:warrant/warrant_bloc.dart';
import 'google_sign_in_screen.dart';

void main() => runApp(const WarrantApp());

class WarrantApp extends StatelessWidget {
  const WarrantApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text('Warrant App'),
        ),
        body: GoogleSignInScreen(), // Display GoogleSignInScreen here
      ),
    );
  }
}
