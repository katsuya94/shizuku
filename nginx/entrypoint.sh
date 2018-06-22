#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# ensure nginx UID is 101
usermod -u 101 nginx

exec "$@"
