#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

exec docker run -it --rm --network shizuku_default \
  --link shizuku_zookeeper_1:zookeeper zookeeper:3.4 \
  zkCli.sh -server zookeeper "$@"
