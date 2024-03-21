//import 'package:billtonapp/google.dart';
import 'package:billtonapp/google.dart';
import 'package:billtonapp/receipt_form.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const ShreeshApp());
}

class ShreeshApp extends StatelessWidget {
  const ShreeshApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Welcome Screen',
      theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(
              seedColor: const Color.fromRGBO(255, 190, 152, 1)),
          useMaterial3: true,
          fontFamily: "Schyler"),
      home: const WelcomeScreen(),
    );
  }
}

class WelcomeScreen extends StatefulWidget {
  const WelcomeScreen({super.key});

  @override
  _WelcomeScreenState createState() => _WelcomeScreenState();
}

class _WelcomeScreenState extends State<WelcomeScreen> {
  int _selectedIndex = 0;
  ValueNotifier<UserCredential?> userCredential = ValueNotifier(null);

  static final List<Widget> _widgetOptions = <Widget>[
    const HomeScreen(),
    const HistoryScreen(),
    const AddPictureScreen(),
    const SettingsScreen(),
    const ProfileScreen(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: userCredential.value != null
              ? CircleAvatar(
                  radius: 20, // Adjust the radius as needed
                  backgroundImage:
                      NetworkImage(userCredential.value!.user!.photoURL!),
                )
              : const Icon(Icons.login),
          onPressed: () async {
            userCredential.value = await signInWithGoogle();
            if (userCredential.value != null) {
              print(userCredential.value!.user!.email);
            }
          },
        ),
        actions: <Widget>[
          IconButton(
            icon: const Icon(Icons.notifications),
            onPressed: () {
              // Navigate to notifications screen
            },
          ),
        ],
      ),
      body: Center(
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        backgroundColor: Theme.of(context).primaryColor, // Set background color
        selectedItemColor:
            Theme.of(context).colorScheme.secondary, // Set selected item color
        unselectedItemColor: Colors.black, // Set unselected item color
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.history),
            label: 'History',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.add_a_photo),
            label: 'Add Picture',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: 'Settings',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
      ),
    );
  }
}

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          Text("Thanks for choosing billton.",
              style: TextStyle(
                  color: Theme.of(context).primaryColor,
                  fontWeight: FontWeight.w500,
                  fontSize: 20.0)),
          const Text("Let's get started!"),
          const Padding(
            padding: EdgeInsets.fromLTRB(20, 100, 20, 0),
            child: Image(image: AssetImage("assets/Welcome.png")),
          ),
        ],
      ),
    );
  }
}

class HistoryScreen extends StatelessWidget {
  const HistoryScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Text('History Screen'),
    );
  }
}

Future<UserCredential?> signInWithGoogle() async {
  await Firebase.initializeApp();

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
    print('Exception->$e');
    return null;
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

class AddPictureScreen extends StatelessWidget {
  const AddPictureScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
        child: Column(
      children: [
        ElevatedButton(
          onPressed: () async {
            File? imageFile = await getImageFromCamera();
            if (imageFile != null) {
              await uploadImageToFirebase(imageFile);
            }
          },
          child: const Text('Take Picture & Upload'),
        ),
        const SizedBox(
          height: 20,
        ),
        ElevatedButton(
          onPressed: () async {
            File? imageFile = await getImageFromGallery();
            if (imageFile != null) {
              await uploadImageToFirebase(imageFile);
            }
          },
          child: const Text('Upload from Gallery'),
        ),
        const SizedBox(
          height: 20,
        ),
        const Padding(
          padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
          child: ReceiptForm(),
        ),
      ],
    ));
  }

  Future<File?> getImageFromGallery() async {
    final picker = ImagePicker();
    final XFile? pickedFile =
        await picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      return File(pickedFile.path);
    } else {
      return null;
    }
  }

  Future<void> uploadImageToFirebase(File imageFile) async {
    try {
      print("uploading");
      String fileName = DateTime.now().millisecondsSinceEpoch.toString();
      Reference ref = FirebaseStorage.instance.ref().child(
          'files/${FirebaseAuth.instance.currentUser!.uid}/$fileName.jpg');
      await ref.putFile(imageFile);
      String imageUrl = await ref.getDownloadURL();
      print('Image uploaded to Firebase: $imageUrl');
    } catch (e) {
      print('Error uploading image to Firebase: $e');
    }
  }

  Future<File?> getImageFromCamera() async {
    final picker = ImagePicker();
    final XFile? pickedFile =
        await picker.pickImage(source: ImageSource.camera);
    if (pickedFile != null) {
      return File(pickedFile.path);
    } else {
      return null;
    }
  }
}

class SettingsScreen extends StatelessWidget {
  const SettingsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Text('Settings Screen'),
    );
  }
}

class ProfileScreen extends StatelessWidget {
  const ProfileScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Text('Profile Screen'),
    );
  }
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const GoogleSignInScreen(),
      debugShowCheckedModeBanner: false,
    );
  }
}
