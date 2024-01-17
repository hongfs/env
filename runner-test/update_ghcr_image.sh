#!/bin/bash

set -e

echo "hongfs----------------------"

pwd

echo "hongfs----------------------"

export

echo "hongfs----------------------"

if [ -f "Dockerfile" ]; then
    sed -i "s+ghcr.io/hongfs/env:+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/env:+g" Dockerfile
fi

if [ -d ".github/workflows/" ]; then
    sed -i "s+ghcr.io/hongfs/env:+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/env:+g" .github/workflows/*
    sed -i "s+ghcr.io/hongfs/ecs-command-actions:main+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/ecs-command-actions:main+g" .github/workflows/*
fi
