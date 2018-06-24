#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

KAFKA_SCRIPT_NAME="$1"
shift

exec docker run -it --rm --network shizuku_default \
  --link shizuku_zookeeper_1:zookeeper \
  --link shizuku_kafka_1:kafka wurstmeister/kafaka:1.1.0 \
  bin/kafka-$KAFKA_SCRIPT_NAME.sh "$@"
