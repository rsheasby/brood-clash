import 'package:brood_clash/backend.dart';

class BcConfiguration extends Configuration {
  BcConfiguration(String configPath) : super.fromFile(File(configPath));

  DatabaseConfiguration database;
}
