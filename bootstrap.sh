#!/usr/bin/env bash
# DigitalOcean Docker 17.12.0~ce on 16.04

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT="$(dirname ${BASH_SOURCE[0]})"

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
  # ensure jenkins user can access jenkins_home
  mkdir -p $PROJECT_ROOT/jenkins_home
  sudo chown 1000 $PROJECT_ROOT/jenkins_home

  mkdir -p $PROJECT_ROOT/nginx/ssl


  # self-sign server certificate
  openssl req -x509 -newkey rsa:2048 -days 365 \
    -keyout $PROJECT_ROOT/nginx/ssl/server.key \
    -out $PROJECT_ROOT/nginx/ssl/server.crt \
    -subj "/C=US/ST=Illinois/L=Chicago/O=atateno.io/OU=atateno.io/CN=atateno.io"

  # restrict server key permissions to the nginx user
  chmod 400 $PROJECT_ROOT/nginx/ssl/server.key
  chown 101:101 $PROJECT_ROOT/nginx/ssl/server.key
fi
