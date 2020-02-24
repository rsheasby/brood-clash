import 'package:aqueduct/aqueduct.dart';

import 'answer.dart';

class Question extends ManagedObject<_Question> implements _Question, Serializable {}

class _Question {
  // TODO This will autoincrement. Apparently this is not desirable so need to
  // find out how to do it another way.
  @primaryKey
  int id;

  String text;

  ManagedSet<Answer> answers;
}
