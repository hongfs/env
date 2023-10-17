#!/bin/bash

set -e

mkdir -p ~/.docker/

alias docker="podman"

cd /runner/actions-runner

echo "128" > /proc/sys/sunrpc/tcp_slot_table_entries

if [ ! -z "$NAS_HOST" ]; then
    mount -t nfs -o vers=4,minorversion=0,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,_netdev,noresvport $NAS_HOST:/cache/ /mnt/cache/
fi

RUNNER_ALLOW_RUNASROOT=1 ./config.sh --url https://github.com/$RUNNER_REPO --token $RUNNER_TOKEN --unattended --disableupdate --ephemeral --labels aliyun
RUNNER_ALLOW_RUNASROOT=1 ./run.sh

sleep 10s
exit 0
