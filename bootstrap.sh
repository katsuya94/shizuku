#!/usr/bin/env bash
# DigitalOcean Docker 17.12.0~ce on 16.04

set -o errexit
set -o nounset
set -o pipefail
set -o verbose

DIRECTORY="$(dirname ${BASH_SOURCE[0]})"
useradd -G sudo,docker -m -s /bin/bash kacchan
passwd kacchan
mkdir -p /home/kacchan/.ssh
cp /root/.ssh/authorized_keys /home/kacchan/.ssh/authorized_keys
chown -R kacchan:kacchan /home/kacchan/.ssh
patch /etc/ssh/sshd_config $DIRECTORY/sshd_config.patch
systemctl restart ssh
