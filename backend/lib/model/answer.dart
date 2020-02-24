import 'package:aqueduct/aqueduct.dart';

import 'question.dart';

class Answer extends ManagedObject<_Answer> implements _Answer, Serializable {}

class _Answer {
  // TODO This will autoincrement. Apparently this is not desirable so need to
  // find out how to do it another way.
  @primaryKey
  int id;

  String text;
  int points;
  bool revealed;

  @Relate(#answers, onDelete: DeleteRule.cascade, isRequired: true)
  Question question;
}
