#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

exec docker run -it --rm --link shizuku_zookeeper:zookeeper zookeeper:3.4 \
  zkCli.sh -server zookeeper "$@"
