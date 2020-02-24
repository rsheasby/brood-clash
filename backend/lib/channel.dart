import 'package:brood_clash/controllers/answer_controller.dart';
import 'package:brood_clash/controllers/question_controller.dart';

import 'bc_config.dart';
import 'brood_clash.dart';
import 'controllers/auth.dart';
import 'controllers/websocket_controller.dart';
import 'services/socket_service.dart';

/// This type initializes an application.
///
/// Override methods in this class to set up routes and initialize services like
/// database connections. See http://aqueduct.io/docs/http/channel/.
class BroodClashChannel extends ApplicationChannel {
  ManagedContext context;
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

    final config = BcConfiguration(options.configurationFilePath);

    final dataModel = ManagedDataModel.fromCurrentMirrorSystem();
    final psc = PostgreSQLPersistentStore.fromConnectionInfo(
        config.database.username,
        config.database.password,
        config.database.host,
        config.database.port,
        config.database.databaseName);

    context = ManagedContext(dataModel, psc);
  }

  /// Construct the request channel.
  ///
  /// Return an instance of some [Controller] that will be the initial receiver
  /// of all [Request]s.
  ///
  /// This method is invoked after [prepare].
  @override
  Controller get entryPoint {
    // TODO: Missing auth.
    return Router()
      ..route("/user")
        .link(() => UserController(context))
      ..route("/questions/[:id]")
        .link(() => QuestionController(context))
      ..route("/answers/:id/reveal")
        .link(() => AnswerController(context))
      ..route("/present")
        .link(() => WebsocketController(socketService))
      ..route("/send")
        .linkFunction((request) async {
          final message = await request.body.decode();
          socketService.broadcast(message);
          messageHub.add(message);
          return Response.ok(null);
        });
  }
}
