#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT="$(dirname $(dirname ${BASH_SOURCE[0]}))"

mkdir -p $PROJECT_ROOT/nginx/ssl

openssl req -newkey rsa:2048 -days 365 -nodes \
  -keyout $PROJECT_ROOT/nginx/ssl/server.key \
  -out $PROJECT_ROOT/nginx/ssl/server.csr \
  -subj "/C=US/ST=Illinois/L=Chicago/O=Adrien Katsuya Tateno/OU=Adrien Katsuya Tateno/CN=atateno.io"
