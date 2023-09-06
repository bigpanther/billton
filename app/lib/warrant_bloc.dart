import 'package:dio/dio.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

sealed class CounterEvent {}

final class CounterIncrementPressed extends CounterEvent {}

final class CounterDecrementPressed extends CounterEvent {}

final class WarrantUserSave extends CounterEvent {
  final String username;
  WarrantUserSave(this.username);
}

final dio = Dio(BaseOptions(baseUrl: 'http://127.0.0.1:8080'));

class WarrantBloc extends Bloc<CounterEvent, int> {
  WarrantBloc() : super(0) {
    on<CounterIncrementPressed>((event, emit) => emit(state + 1));
    on<CounterDecrementPressed>((event, emit) => emit(state - 1));
    on<WarrantUserSave>((event, emit) async {
      dio.options.contentType = Headers.jsonContentType;
      var r =
          await dio.post('/users', data: {'name': '${event.username}-$state'});
      emit(state + 3);
    });
  }
}
