import 'dart:async';
import 'dart:io';

import 'package:aqueduct/aqueduct.dart';
import 'socket_service.dart';

class WebsocketController extends Controller {
  WebsocketController(this.socketService);

  final SocketService socketService;

  @override
  FutureOr<RequestOrResponse> handle(Request request) async {
    final socket = await WebSocketTransformer.upgrade(request.raw);
    socket.listen((_) {}, onDone: socket.close);
    socketService.register(socket);
    return null;
  }
}