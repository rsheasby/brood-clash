import 'brood_clash.dart';
import 'socket_service.dart';
import 'websocket_controller.dart';

/// This type initializes an application.
///
/// Override methods in this class to set up routes and initialize services like
/// database connections. See http://aqueduct.io/docs/http/channel/.
class BroodClashChannel extends ApplicationChannel {
  SocketService socketService;

  /// Initialize services in this method.
  ///
  /// Implement this method to initialize services, read values from [options]
  /// and any other initialization required before constructing [entryPoint].
  ///
  /// This method is invoked prior to [entryPoint] being accessed.
  @override
  Future prepare() async {
    logger.onRecord.listen(
        (rec) => print("$rec ${rec.error ?? ""} ${rec.stackTrace ?? ""}"));

    socketService = SocketService();

    messageHub.listen((data) {
      socketService.broadcast(data);
    });
  }

  /// Construct the request channel.
  ///
  /// Return an instance of some [Controller] that will be the initial receiver
  /// of all [Request]s.
  ///
  /// This method is invoked after [prepare].
  @override
  Controller get entryPoint {
    final router = Router();

    router.route("/present").link(() => WebsocketController(socketService));
    router.route("/send").linkFunction((request) async {
      final message = await request.body.decode();
      socketService.broadcast(message);
      messageHub.add(message);
      return Response.ok(null);
    });

    return router;
  }
}
