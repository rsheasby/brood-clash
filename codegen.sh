#!/bin/bash

FILENAME='swagger-codegen-cli.jar'
URL='http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.10/swagger-codegen-cli-2.4.10.jar'
OUTPUT_DIR='swagger-client'
SWAGGER_FILE='backend/swagger.yml'

if [ ! -e $FILENAME ]; then
	echo $FILENAME "does not exist, downloading..."
	wget -nv --show-progress $URL -O $FILENAME
else
	echo $FILENAME "exists."
fi

java -jar swagger-codegen-cli.jar generate \
	-i $SWAGGER_FILE                       \
	-l javascript                          \
	-DuseES6=true                          \
	-DusePromises=true                     \
	-o $OUTPUT_DIR
