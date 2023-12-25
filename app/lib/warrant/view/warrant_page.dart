import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../warrant_bloc.dart';

class WarrantPage extends StatelessWidget {
  const WarrantPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Warrant')),
      body: BlocBuilder<WarrantBloc, int>(
        builder: (context, count) {
          return Center(
            child: Text(
              '$count',
              style: const TextStyle(fontSize: 24.0),
            ),
          );
        },
      ),
      floatingActionButton: Column(
        crossAxisAlignment: CrossAxisAlignment.end,
        mainAxisAlignment: MainAxisAlignment.end,
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 5.0),
            child: FloatingActionButton(
              child: const Icon(Icons.save),
              onPressed: () =>
                  context.read<WarrantBloc>().add(WarrantUserSave("harsimran")),
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 5.0),
            child: FloatingActionButton(
              child: const Icon(Icons.add),
              onPressed: () =>
                  context.read<WarrantBloc>().add(CounterIncrementPressed()),
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 5.0),
            child: FloatingActionButton(
              child: const Icon(Icons.remove),
              onPressed: () =>
                  context.read<WarrantBloc>().add(CounterDecrementPressed()),
            ),
          ),
        ],
      ),
    );
  }
}
