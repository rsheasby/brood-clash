import 'dart:io';

import 'package:aqueduct/aqueduct.dart';

class BcConfiguration extends Configuration {
  BcConfiguration(String configPath) : super.fromFile(File(configPath));

  DatabaseConfiguration database;
}
