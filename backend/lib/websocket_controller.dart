import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:aqueduct/aqueduct.dart';

class WebsocketController extends Controller {
  @override
  FutureOr<RequestOrResponse> handle(Request request) async {
    // TODO: Never closed.
    final socket = await WebSocketTransformer.upgrade(request.raw);
    final content = jsonEncode({
      "message": "hello",
    });
    socket.add(content);
    return Response(HttpStatus.switchingProtocols, null, null);
  }
}