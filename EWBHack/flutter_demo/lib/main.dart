import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Moral Code Hackathon',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // TRY THIS: Try running your application with "flutter run". You'll see
        // the application has a teal toolbar. Then, without quitting the app,
        // try changing the seedColor in the colorScheme below to Colors.green
        // and then invoke "hot reload" (save your changes or press the "hot
        // reload" button in a Flutter-supported IDE, or press "r" if you used
        // the command line to start the app).
        //
        // Notice that the counter didn't reset back to zero; the application
        // state is not lost during the reload. To reset the state, use hot
        // restart instead.
        //
        // This works for code too, not just values: Most code changes can be
        // tested with just a hot reload.
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.tealAccent),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Engineers Without Borders'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;
  String _recipetext = "Please input your data…";

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }

  void _updateResponse(String text) {
    setState(() {
      _recipetext = text;
    });
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      appBar: AppBar(
        // TRY THIS: Try changing the color here to a specific color (to
        // Colors.amber, perhaps?) and trigger a hot reload to see the AppBar
        // change color while the other colors stay the same.
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text(widget.title),
      ),
      body: Center(
        // Center is a layout widget. It takes a single child and positions it
        // in the middle of the parent.
        child: Column(
          // Column is also a layout widget. It takes a list of children and
          // arranges them vertically. By default, it sizes itself to fit its
          // children horizontally, and tries to be as tall as its parent.
          //
          // Column has various properties to control how it sizes itself and
          // how it positions its children. Here we use mainAxisAlignment to
          // center the children vertically; the main axis here is the vertical
          // axis because Columns are vertical (the cross axis would be
          // horizontal).
          //
          // TRY THIS: Invoke "debug painting" (choose the "Toggle Debug Paint"
          // action in the IDE, or press "p" in the console), to see the
          // wireframe for each widget.
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Column(
              children: [
                /* Starter code */
                const Text(
                  'You have clicked the demo button this many times:',
                ),
                /* Form field */
                FormExample(
                  onChanged: _updateResponse,
                ),
                /* Buttons */
                const Text('What is your level of exercise?'),
                ElevatedButton(
                    onPressed: () {
                      print("Button pressed!");
                      _updateResponse("Please input your data…");
                    },
                    child: const Text('Button')),
              ],
            ),
            Column(children: [
              Text(
                '$_recipetext',
                style: Theme.of(context).textTheme.headlineMedium,
              ),
            ],)
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}

/* Form auxiliries */
enum ColorLabel {
  blue('Male', Colors.blue),
  pink('Female', Colors.pink),
  green('Non-binary', Colors.green);

  const ColorLabel(this.label, this.color);
  final String label;
  final Color color;
}

void geminiAPIcall(Map<String,String> jsonstr, ValueChanged<String> onChanged) async {
  final response = await http.post(
    Uri.parse('http://localhost:8080/geminiResp'),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(jsonstr),
  );

  print(response.statusCode);

  if (response.statusCode < 300) {
    print(response.body);
    onChanged(response.body);
  } else {}
}

/* Form */
class FormExample extends StatefulWidget {
  const FormExample({super.key, required this.onChanged});
  final ValueChanged<String> onChanged;

  @override
  State<FormExample> createState() => _FormExampleState();
}

class _FormExampleState extends State<FormExample> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController _genderController = TextEditingController();
  final TextEditingController _ageController = TextEditingController();
  final _calorieController = <TextEditingController>[TextEditingController(),TextEditingController()];
  final _percentController = <TextEditingController>[TextEditingController(),TextEditingController(),TextEditingController()];
  final TextEditingController _dietaryController = TextEditingController();
  final TextEditingController _allergyController = TextEditingController();
  ColorLabel? selectedColor;

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          /* Row 1 */
          Row(
            children: [
              SizedBox(width: 8),
              /* Gender */
              Expanded(
                  flex: 3,
                  child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        DropdownMenu<ColorLabel>(
                          initialSelection: ColorLabel.green,
                          controller: _genderController,
                          requestFocusOnTap: true,
                          label: const Text('Color'),
                          onSelected: (ColorLabel? color) {
                            setState(() {
                              selectedColor = color;
                            });
                          },
                          dropdownMenuEntries: ColorLabel.values
                              .map<DropdownMenuEntry<ColorLabel>>(
                                  (ColorLabel color) {
                            return DropdownMenuEntry<ColorLabel>(
                              value: color,
                              label: color.label,
                              enabled: color.label != 'Grey',
                              style: MenuItemButton.styleFrom(
                                foregroundColor: color.color,
                              ),
                            );
                          }).toList(),
                        ),
                      ])),
              SizedBox(width: 8),
              /* Age */
              Expanded(
                flex: 3,
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextFormField(
                        controller: _ageController,
                        decoration: const InputDecoration(
                          hintText: 'Age',
                        ),
                        validator: (String? value) {
                          if (value == null ||
                              value.isEmpty ||
                              (num.tryParse(value) == null
                                      ? -1
                                      : num.parse(value)) <
                                  0) {
                            return 'Enter a nonnegative number';
                          }
                          return null;
                        },
                      ),
                    ]),
              ),
              SizedBox(width: 8),
            ],
          ),

          /* Row 2: Calories */
          Row(
            children: [
              SizedBox(width: 8),
              Expanded(
                  flex: 3,
                  child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        /* Caloric intake */
                        TextFormField(
                          controller: _calorieController[0],
                          decoration: const InputDecoration(
                            hintText: 'Total calories',
                          ),
                          validator: (String? value) {
                            if (value == null ||
                                value.isEmpty ||
                                (num.tryParse(value) == null
                                        ? -1
                                        : num.parse(value)) <=
                                    0) {
                              return 'Enter a positive number';
                            }
                            return null;
                          },
                        ),
                      ])),
              SizedBox(width: 8),
              Expanded(
                flex: 3,
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      /* Dietary restrictions */
                      TextFormField(
                        controller: _calorieController[1],
                        decoration: const InputDecoration(
                          hintText: 'Daily caloric consumption',
                        ),
                        validator: (String? value) {
                          if (value == null ||
                              value.isEmpty ||
                              (num.tryParse(value) == null
                                      ? -1
                                      : num.parse(value)) <=
                                  0) {
                            return 'Enter a positive number';
                          }
                          return null;
                        },
                      ),
                    ]),
              ),
              SizedBox(width: 8),
            ],
          ),

          /* Row 3: Caloric percentage */
          Row(
            children: [
              SizedBox(width: 8),
              Expanded(
                  flex: 3,
                  child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        TextFormField(
                          controller: _percentController[0],
                          decoration: const InputDecoration(
                            hintText: 'Percentage carbs',
                          ),
                          validator: (String? value) {
                            if (value == null || value.isEmpty) {
                              return 'Carb + Protein + Fat == 100';
                            }
                            return null;
                          },
                        ),
                      ])),
              SizedBox(width: 8),
              Expanded(
                flex: 3,
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextFormField(
                        controller: _percentController[1],
                        decoration: const InputDecoration(
                          hintText: 'Percentage fat',
                        ),
                        validator: (String? value) {
                          if (value == null || value.isEmpty) {
                            return 'Carb + Protein + Fat == 100';
                          }
                          return null;
                        },
                      ),
                    ]),
              ),
              SizedBox(width: 8),
              Expanded(
                flex: 3,
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextFormField(
                        controller: _percentController[2],
                        decoration: const InputDecoration(
                          hintText: 'Percentage protein',
                        ),
                        validator: (String? value) {
                          if (value == null || value.isEmpty) {
                            return 'Carb + Protein + Fat == 100';
                          }
                          return null;
                        },
                      ),
                    ]),
              ),
              SizedBox(width: 8),
            ],
          ),

          /* Lower rows: Dietary restrictions */
          TextFormField(
            controller: _dietaryController,
            decoration: const InputDecoration(
              hintText: 'Enter your dietary restriction',
            ),
            validator: (String? value) {
              if (value == null || value.isEmpty) {
                return 'lacto-ovo, vegetarian, diabetic, halal, …';
              }
              return null;
            },
          ),
          TextFormField(
            controller: _allergyController,
            decoration: const InputDecoration(
              hintText: 'Enter your allergies',
            ),
            validator: (String? value) {
              if (value == null || value.isEmpty) {
                return 'peanuts, gluten, …';
              }
              return null;
            },
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: ElevatedButton(
              onPressed: () {
                // Validate will return true if the form is valid, or false if
                // the form is invalid.
                if (_formKey.currentState!.validate()) {
                  // Process data.
                  print(_genderController.value);
                  print(_ageController.value);
                  // Perform API call
                  final body = <String,String> {
                    'gender': _genderController.value.toString(),
                    'age': _ageController.value.toString(),
                    'bmi': "20",
                    'calories': _calorieController[0].value.toString(),
                    'ETotal': _calorieController[1].value.toString(),
                    'carbs': _percentController[0].value.toString(),
                    'fat': _percentController[1].value.toString(),
                    'protein': _percentController[2].value.toString(),
                    'diet': _dietaryController.value.toString(),
                    'allergy': _allergyController.value.toString()
                  };
                  geminiAPIcall(body, widget.onChanged);
                }
              },
              child: const Text('Submit'),
            ),
          ),
        ],
      ),
    );
  }
}
