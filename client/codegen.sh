#!/bin/bash

CMD='openapi-generator'
OUTPUT_DIR='api/codegen'
SWAGGER_FILE='../backend/docs/swagger.yaml'

$CMD generate \
   --skip-validate-spec \
	-i $SWAGGER_FILE                       \
	-g typescript-axios                    \
	-o $OUTPUT_DIR
