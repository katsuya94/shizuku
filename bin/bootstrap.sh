#!/usr/bin/env bash
# DigitalOcean Docker 17.12.0~ce on 16.04

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT="$(dirname $(dirname ${BASH_SOURCE[0]}))"

if [ $(id -u) = 0 ]; then
  # create admin user
  useradd -G sudo,docker -m -s /bin/bash kacchan

  # set admin user password
  passwd kacchan

  # copy authorized_keys to admin user
  mkdir -p /home/kacchan/.ssh
  cp /root/.ssh/authorized_keys /home/kacchan/.ssh/authorized_keys
  chown -R kacchan:kacchan /home/kacchan/.ssh

  # configure SSH daemon
  patch /etc/ssh/sshd_config $PROJECT_ROOT/sshd_config.patch
  systemctl restart ssh

  # install packages
  apt install htop
else
  # restrict server key permissions to the nginx user
  sudo chown 101:101 $PROJECT_ROOT/nginx/ssl/server.key
  sudo chmod 400 $PROJECT_ROOT/nginx/ssl/server.key

  # restrict gmail API credential permissions
  sudo chown 0:0 $PROJECT_ROOT/mailstream/config/client_secret.json
  sudo chmod 400 $PROJECT_ROOT/mailstream/config/client_secret.json
fi
