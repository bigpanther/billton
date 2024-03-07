import 'package:billtonapp/google.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:image_picker/image_picker.dart';
import 'package:firebase_storage/firebase_storage.dart' as firebase_storage;
import 'dart:io';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  // await Firebase.initializeApp();
  runApp(ShreeshApp());
}

class ShreeshApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Welcome Screen',
      theme: ThemeData(
          colorScheme:
              ColorScheme.fromSeed(seedColor: Color.fromRGBO(255, 190, 152, 1)),
          useMaterial3: true,
          fontFamily: "Schyler"),
      home: WelcomeScreen(),
    );
  }
}

class WelcomeScreen extends StatefulWidget {
  @override
  _WelcomeScreenState createState() => _WelcomeScreenState();
}

class _WelcomeScreenState extends State<WelcomeScreen> {
  int _selectedIndex = 0;
  ValueNotifier<UserCredential?> userCredential = ValueNotifier(null);

  static List<Widget> _widgetOptions = <Widget>[
    HomeScreen(),
    HistoryScreen(),
    AddPictureScreen(),
    SettingsScreen(),
    ProfileScreen(),
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
              ? Image.network(userCredential.value!.user!.photoURL!)
              : Icon(Icons.login),
          onPressed: () async {
            userCredential.value = await signInWithGoogle();
            if (userCredential.value != null) {
              print(userCredential.value!.user!.email);
            }
          },
        ),
        actions: <Widget>[
          IconButton(
            icon: Icon(Icons.notifications),
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
          Text("Let's get started!"),
          Padding(
            padding: EdgeInsets.fromLTRB(20, 100, 20, 0),
            child: Image(image: AssetImage("assets/Welcome.png")),
          ),
        ],
      ),
    );
  }
}

class HistoryScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Center(
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
      ],
    ));
  }

  Future<File?> getImageFromGallery() async {
    final picker = ImagePicker();
    final PickedFile? pickedFile =
        await picker.getImage(source: ImageSource.gallery);
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
      firebase_storage.Reference ref = firebase_storage.FirebaseStorage.instance
          .ref()
          .child(
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
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Text('Settings Screen'),
    );
  }
}

class ProfileScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Center(
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
