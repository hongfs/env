#!/bin/bash

set -e

if [ "$HAS_HONGKONG_VPC_REGISTRY" != "1" ]; then
    exit 0
fi

if [ -e "$(pwd)/Dockerfile" ]; then
    echo "Dockerfile exists"

    sed -i "s+ghcr.io/hongfs/env:+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/env:+g" Dockerfile
fi

if [ -d ".github/workflows/" ]; then
    echo ".github/workflows/ exists"

    sed -i "s+ghcr.io/hongfs/env:+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/env:+g" .github/workflows/*
    sed -i "s+ghcr.io/hongfs/ecs-command-actions:main+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/ecs-command-actions:main+g" .github/workflows/*
fi
