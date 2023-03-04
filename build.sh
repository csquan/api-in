#!/bin/bash

set -eu

tag="${1}"
echo "tag is ${tag}"

docker build --pull --force-rm --no-cache -t "reg.huiwang.io/fat/coin-manage:${tag}" .
digest=$(docker push "reg.huiwang.io/fat/coin-manage:${tag}" | awk '/digest/{print $3}')
cosign sign --key ~/cosign.key "reg.huiwang.io/fat/coin-manage@${digest}"
