#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

cat > /var/jenkins_home/log.properties <<EOF
handlers=java.util.logging.ConsoleHandler
jenkins.level=FINEST
java.util.logging.ConsoleHandler.level=FINEST
EOF

exec "$@"
