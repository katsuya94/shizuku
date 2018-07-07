#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

KAFKA_SCRIPT_NAME="$1"
shift

exec docker run -it --rm --network shizuku_default wurstmeister/kafka:1.1.0 \
  kafka-$KAFKA_SCRIPT_NAME.sh "$@"
