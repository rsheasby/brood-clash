import 'package:brood_clash/brood_clash.dart';
import 'package:brood_clash/model/question.dart';

class QuestionController extends ResourceController {
  QuestionController(this.context);

  final ManagedContext context;

  @Operation.get()
  Future<Response> getQuestions() async {
    final query = Query<Question>(context)..join(set: (q) => q.answers);
    final questions = await query.fetch();
    return Response.ok(questions);
  }

  @Operation.get("id")
  Future<Response> getQuestion(@Bind.path("id") int id) async {
    final query = Query<Question>(context)
      ..join(set: (q) => q.answers)
      ..where((q) => q.id).equalTo(id);
    final question = await query.fetchOne();
    return Response.ok(question);
  }

  @Operation.post()
  Future<Response> createQuestions(
      @Bind.body() List<Question> questions) async {
    // TODO: No idea how to do this properly. Using context.insertObjects does
    // not create answer entities.
    final insertedIds = await context.transaction((transaction) async {
      final List<int> insertedObjects = [];
      for (final question in questions) {
        final insertedQuestion = await transaction.insertObject(question);
        insertedObjects.add(insertedQuestion.id);
        for (final answer in question.answers) {
          answer.question = insertedQuestion;
          await transaction.insertObject(answer);
        }
      }
      return insertedObjects;
    });

    final query = Query<Question>(context)
      ..where((q) => q.id).oneOf(insertedIds)
      ..join(set: (q) => q.answers);

    return Response.ok(await query.fetch());
  }
}
