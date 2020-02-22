import 'package:brood_clash/backend.dart';
import 'package:brood_clash/model/user.dart';

class UserController extends ResourceController {
  final ManagedContext context;

  UserController(this.context);

  @Operation.post()
  Future<Response> createUser(@Bind.body() User user) async {
    final insertedUser = await context.insertObject(user);

    return Response.ok(insertedUser);
  }
}