import 'package:flutter/material.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';
import 'package:form_builder_validators/form_builder_validators.dart';

class ReceiptForm extends StatefulWidget {
  const ReceiptForm({super.key});

  @override
  State<ReceiptForm> createState() => _ReceiptFormState();
}

class _ReceiptFormState extends State<ReceiptForm> {
  final _formKey = GlobalKey<FormBuilderState>();
  final _emailFieldKey = GlobalKey<FormBuilderFieldState>();

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: FormBuilder(
        key: _formKey,
        child: Column(
          children: [
            FormBuilderTextField(
              name: 'Store_name',
              decoration: const InputDecoration(labelText: 'Store Name'),
              validator: FormBuilderValidators.compose([
                FormBuilderValidators.required(),
              ]),
            ),
            const SizedBox(height: 10),
            FormBuilderTextField(
              key: _emailFieldKey,
              name: 'brand_name',
              decoration: const InputDecoration(labelText: 'Brand Name'),
              validator: FormBuilderValidators.compose([
                FormBuilderValidators.required(),
                FormBuilderValidators.email(),
              ]),
            ),
            const SizedBox(height: 10),
            FormBuilderDateTimePicker(
              name: 'transaction_time',
              decoration: InputDecoration(labelText: 'Transaction Time'),
              initialValue: DateTime.now(), // Initial value (optional)
              onChanged: (value) {
                print(value);
              },
            ),
            const SizedBox(height: 10),
            FormBuilderDateTimePicker(
              name: 'expiry_date',
              inputType: InputType.date,
              decoration: InputDecoration(labelText: 'Expiry Date'),
              initialValue: DateTime.now(), // Initial value (optional)
              onChanged: (value) {
                // Handle the selected expiry date
                print(value);
              },
            ),
            const SizedBox(height: 10),
            FormBuilderTextField(
              name: 'amount',
              decoration: const InputDecoration(labelText: 'Amount (CAD)'),
              validator: FormBuilderValidators.compose([
                FormBuilderValidators.required(),
              ]),
            ),
            const SizedBox(height: 10),
            MaterialButton(
              color: Theme.of(context).colorScheme.secondary,
              onPressed: () {
                if (_formKey.currentState?.saveAndValidate() ?? false) {
                  if (true) {
                    // Either invalidate using Form Key
                    _formKey.currentState?.fields['email']
                        ?.invalidate('Email already taken.');
                    // OR invalidate using Field Key
                    // _emailFieldKey.currentState?.invalidate('Email already taken.');
                  }
                }
                debugPrint(_formKey.currentState?.value.toString());
              },
              child: const Text('Add Receipt',
                  style: TextStyle(color: Colors.white)),
            )
          ],
        ),
      ),
    );
  }
}
