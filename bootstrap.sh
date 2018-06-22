#!/usr/bin/env bash
# DigitalOcean Docker 17.12.0~ce on 16.04
set -o errexit
set -o verbose
set -o nounset
set -o pipefail

useradd -G sudo, docker kacchan
mkdir -p /home/kacchan/.ssh
cp /root/.ssh/authorized_keys /home/kacchan/.ssh/authorized_keys
su -l kacchan passwd