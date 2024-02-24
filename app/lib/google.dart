// import 'package:google_sign_in/google_sign_in.dart';

// class GoogleSignInScreen extends StatefulWidget {
//   const GoogleSignInScreen({Key? key}) : super(key: key);

//   @override
//   State<GoogleSignInScreen> createState() => _GoogleSignInScreenState();
// }

// class _GoogleSignInScreenState extends State<GoogleSignInScreen> {
//   ValueNotifier userCredential = ValueNotifier('');

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//         appBar: AppBar(title: const Text('Google SignIn Screen')),
//         body: ValueListenableBuilder(
//             valueListenable: userCredential,
//             builder: (context, value, child) {
//               return (userCredential.value == '' ||
//                       userCredential.value == null)
//                   ? Center(
//                       child: Card(
//                         elevation: 5,
//                         shape: RoundedRectangleBorder(
//                             borderRadius: BorderRadius.circular(10)),
//                         child: IconButton(
//                           iconSize: 40,
//                           icon: Image.asset(
//                             'assets/images/google_icon.png',
//                           ),
//                           onPressed: () async {
//                             userCredential.value = await signInWithGoogle();
//                             if (userCredential.value != null)
//                               print(userCredential.value.user!.email);
//                           },
//                         ),
//                       ),
//                     )
//                   : Center(
//                       child: Column(
//                         crossAxisAlignment: CrossAxisAlignment.center,
//                         mainAxisAlignment: MainAxisAlignment.center,
//                         children: [
//                           Container(
//                             clipBehavior: Clip.antiAlias,
//                             decoration: BoxDecoration(
//                                 shape: BoxShape.circle,
//                                 border: Border.all(
//                                     width: 1.5, color: Colors.black54)),
//                             child: Image.network(
//                                 userCredential.value.user!.photoURL.toString()),
//                           ),
//                           const SizedBox(
//                             height: 20,
//                           ),
//                           Text(userCredential.value.user!.displayName
//                               .toString()),
//                           const SizedBox(
//                             height: 20,
//                           ),
//                           Text(userCredential.value.user!.email.toString()),
//                           const SizedBox(
//                             height: 30,
//                           ),
//                           ElevatedButton(
//                               onPressed: () async {
//                                 bool result = await signOutFromGoogle();
//                                 if (result) userCredential.value = '';
//                               },
//                               child: const Text('Logout'))
//                         ],
//                       ),
//                     );
//             }));
//   }


// Future<dynamic> signInWithGoogle() async {
//     try {
//       final GoogleSignInAccount? googleUser = await GoogleSignIn().signIn();

//       final GoogleSignInAuthentication? googleAuth =
//           await googleUser?.authentication;

//       final credential = GoogleAuthProvider.credential(
//         accessToken: googleAuth?.accessToken,
//         idToken: googleAuth?.idToken,
//       );

//       return await FirebaseAuth.instance.signInWithCredential(credential);
//     } on Exception catch (e) {
//       // TODO
//       print('exception->$e');
//     }
//   }

// Future<bool> signOutFromGoogle() async {
//     try {
//       await FirebaseAuth.instance.signOut();
//       return true;
//     } on Exception catch (_) {
//       return false;
//     }
//   }
import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:camera/camera.dart';
import 'camera_screen.dart'; // Import the CameraScreen widget

class GoogleSignInScreen extends StatefulWidget {
  const GoogleSignInScreen({Key? key}) : super(key: key);

  @override
  State<GoogleSignInScreen> createState() => _GoogleSignInScreenState();
}

class _GoogleSignInScreenState extends State<GoogleSignInScreen> {
  ValueNotifier userCredential = ValueNotifier('');

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Google SignIn Screen'),
      ),
      body: ValueListenableBuilder(
        valueListenable: userCredential,
        builder: (context, value, child) {
          return (userCredential.value == '' || userCredential.value == null)
              ? Center(
                  child: Card(
                    elevation: 5,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                    child: IconButton(
                      iconSize: 40,
                      icon: Image.asset(
                        'assets/images/google_icon.png',
                      ),
                      onPressed: () async {
                        userCredential.value = await signInWithGoogle();
                        if (userCredential.value != null) {
                          print(userCredential.value.user!.email);
                        }
                      },
                    ),
                  ),
                )
              : Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    Container(
                      clipBehavior: Clip.antiAlias,
                      decoration: BoxDecoration(
                        shape: BoxShape.circle,
                        border: Border.all(
                          width: 1.5,
                          color: Colors.black54,
                        ),
                      ),
                      child: Image.network(
                        userCredential.value.user!.photoURL.toString(),
                      ),
                    ),
                    const SizedBox(height: 20),
                    Text(
                      userCredential.value.user!.displayName.toString(),
                    ),
                    const SizedBox(height: 20),
                    Text(
                      userCredential.value.user!.email.toString(),
                    ),
                    const SizedBox(height: 30),
                    ElevatedButton(
                      onPressed: () async {
                        bool result = await signOutFromGoogle();
                        if (result) userCredential.value = '';
                      },
                      child: const Text('Logout'),
                    ),
                    const SizedBox(height: 20),
                    ElevatedButton(
                      onPressed: () async {
                        await _initializeCameraAndNavigate(context);
                      },
                      style: ElevatedButton.styleFrom(
                        primary: Colors.black,
                        onPrimary: Colors.white,
                        padding: EdgeInsets.all(16),
                        minimumSize: Size.square(64),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(0),
                        ),
                        textStyle: TextStyle(fontSize: 18),
                      ),
                      child: const Text('Take a Picture'),
                    ),
                  ],
                );
        },
      ),
    );
  }

  Future<dynamic> signInWithGoogle() async {
    try {
      final GoogleSignInAccount? googleUser = await GoogleSignIn().signIn();

      final GoogleSignInAuthentication? googleAuth =
          await googleUser?.authentication;

      final credential = GoogleAuthProvider.credential(
        accessToken: googleAuth?.accessToken,
        idToken: googleAuth?.idToken,
      );

      return await FirebaseAuth.instance.signInWithCredential(credential);
    } on Exception catch (e) {
      // TODO
      print('exception->$e');
    }
  }

  Future<bool> signOutFromGoogle() async {
    try {
      await FirebaseAuth.instance.signOut();
      return true;
    } on Exception catch (_) {
      return false;
    }
  }

  Future<void> _initializeCameraAndNavigate(BuildContext context) async {
    final cameras = await availableCameras();
    final firstCamera = cameras.first;

    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => CameraScreen(camera: firstCamera),
      ),
    );
  }
}
