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
  runApp(const BilltonApp());
}

class BilltonApp extends StatelessWidget {
  const BilltonApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Welcome Screen',
      theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(
              seedColor: const Color(0xFF388E3C),),
          useMaterial3: true,
          fontFamily: "PlayfairDisplay-VariableFont_wght"),
      home: const WelcomeScreen(),
    );
  }
}

class WelcomeScreen extends StatefulWidget {
  const WelcomeScreen({super.key});

  @override
  // ignore: library_private_types_in_public_api
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
     
     // toolbarHeight: 60,
        leading:   Row(
          children: [
            Image.asset(
              'assets/logo.png',
              width: 50,
              height: 50,
            ),
            const SizedBox(width: 4),
          ],),
          title: const Center(
            child: Text(
                "BillTon",
              style: TextStyle(
                color: Color(0xFF388E3C), // Adjust the text color to make it visible
                fontWeight: FontWeight.bold,
                fontSize: 25.0,
              ),
            ),
          
        ),


        actions: <Widget>[
          Row(
            children: [
              IconButton(
                icon: const Icon(Icons.notifications),
                onPressed: () {
                  // Navigate to notifications screen
                },
              ),
              IconButton(
                icon: userCredential.value != null
                    ? CircleAvatar(
                        radius: 20, // Adjust the radius as needed
                        backgroundImage: NetworkImage(userCredential.value!.user!.photoURL!),
                      )
                    : const Icon(Icons.login),
                onPressed: () async {
                  userCredential.value = await signInWithGoogle();
                  if (userCredential.value != null) {
                    print(userCredential.value!.user!.email);
                  }
                },
              ),
              
            ],
          ),
        ],
      ),

       body: Stack(
        children: [
          Container(
            decoration: const BoxDecoration(
              image: DecorationImage(
                image: AssetImage("assets/welcome_background.png"), // Your background image
                fit: BoxFit.cover,
              ),
            ),
          ),
          Center(
            child: _widgetOptions.elementAt(_selectedIndex),
          ),
        ],
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
          Padding(padding:const EdgeInsets.only(top: 80.0),
                  child:Text("Warranty in your pocket",
                    style: TextStyle(
                    color: Theme.of(context).primaryColor,
                   
                    fontSize: 30.0,
                    fontStyle: FontStyle.italic,
                    fontFamily: "PlayfairDisplay-Italic-VariableFont_wght",)),
          ),
          const Padding(
            padding: EdgeInsets.fromLTRB(20, 10, 20, 0),
            child: Image(image: AssetImage("assets/welcome 2.png")),
          ),
          const     Text("Welcome,",
                    style: TextStyle(
                    color: Color(0xFF388E3C),
                    fontWeight: FontWeight.bold,
                    fontSize: 30.0,
                    fontFamily: "PlayfairDisplay-Italic-VariableFont_wght",)),
          const     Text("please sign in.",
                    style: TextStyle(
                    color: Color(0xFF388E3C),
                    fontSize: 20.0,
                    fontFamily: "PlayfairDisplay-Italic-VariableFont_wght",)),
          const SizedBox(height: 20),
          MaterialButton(
                onPressed: () async {
                  UserCredential? userCredential = await signInWithGoogle();
                  if (userCredential != null) {
                    // Assuming you want to do something with the email or navigate after login
                    print(userCredential.user!.email);  // Debug: Print email to console
                    // You might want to navigate to a new screen or update state here
                    Navigator.push(
                      // ignore: use_build_context_synchronously
                      context,
                      MaterialPageRoute(builder: (context) => const AddPictureScreen()), // Replace NextScreen with your actual next screen
                    );
                  } else {
                    // Login failed or was cancelled, handle accordingly
                    // ignore: use_build_context_synchronously
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text("Login failed, please try again."))
                    );
                  }
                },
                color: Theme.of(context).colorScheme.secondary,
                padding: EdgeInsets.zero,
                child: Image.asset(
                  'assets/GoogleSignUp.png', // Ensure this path is correct
                  height: 40.0,
                ),
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
