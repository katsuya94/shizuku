#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT="$(dirname $(dirname ${BASH_SOURCE[0]}))"
ARCHIVE_PATH="/root/jenkins_home-$(date "+%Y%m%d%H%M%S").tar.gz"

docker-compose -f $PROJECT_ROOT/docker-compose.yaml stop jenkins
docker run --rm --volumes-from shizuku_jenkins_1 -v ~:/root ubuntu \
  tar cvzf $ARCHIVE_PATH -C /var jenkins_home
sudo chown $(id -u):$(id -g) $ARCHIVE_PATH
docker-compose -f $PROJECT_ROOT/docker-compose.yaml start jenkins
