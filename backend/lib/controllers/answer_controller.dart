import 'package:aqueduct/aqueduct.dart';
import 'package:brood_clash/model/answer.dart';

class AnswerController extends ResourceController {
  AnswerController(this.context);

  final ManagedContext context;

  @Operation.post("id")
  Future<Response> reveal(@Bind.path("id") int id) async {
    final query = Query<Answer>(context)
        ..where((a) => a.id).equalTo(id)
        ..values.revealed = true;
    await query.updateOne();
    return Response.ok(null);
  }
}
