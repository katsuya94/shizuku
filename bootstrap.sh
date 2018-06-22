#!/usr/bin/env bash
# DigitalOcean Docker 17.12.0~ce on 16.04

set -o errexit
set -o nounset
set -o pipefail
set -o verbose

PROJECT_ROOT="$(dirname ${BASH_SOURCE[0]})"

if [ $(id -u) = 0 ]; then
  useradd -G sudo,docker -m -s /bin/bash kacchan
  passwd kacchan
  mkdir -p /home/kacchan/.ssh
  cp /root/.ssh/authorized_keys /home/kacchan/.ssh/authorized_keys
  chown -R kacchan:kacchan /home/kacchan/.ssh
  patch /etc/ssh/sshd_config $PROJECT_ROOT/sshd_config.patch
  systemctl restart ssh
else
  mkdir -p $PROJECT_ROOT/jenkins_home
  sudo chown 1000 $PROJECT_ROOT/jenkins_home
fi
