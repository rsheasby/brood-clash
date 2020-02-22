import 'dart:async';

import 'package:aqueduct/aqueduct.dart';
import 'package:brood_clash/model/user.dart';

class UserController extends ResourceController {
  UserController(this.context);

  final ManagedContext context;

  @Operation.post()
  Future<Response> createUser(@Bind.body() User user) async {
    final insertedUser = await context.insertObject(user);

    return Response.ok(insertedUser);
  }
}