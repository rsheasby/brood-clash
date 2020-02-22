import 'package:brood_clash/backend.dart';

class User extends ManagedObject<_User> implements _User , Serializable {}

class _User {
  @Column(primaryKey: true)
  int id;

  String username;
  String code;
}