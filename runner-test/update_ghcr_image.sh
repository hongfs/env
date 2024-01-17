#!/bin/bash

set -e

dir=$(pwd)

sed -i "s+ghcr.io/hongfs/env:+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/env:+g" $(dir)/Dockerfile

if [ -d ".github/workflows/" ]; then
    echo ".github/workflows/ exists"

    sed -i "s+ghcr.io/hongfs/env:+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/env:+g" .github/workflows/*
    sed -i "s+ghcr.io/hongfs/ecs-command-actions:main+registry-vpc.cn-hongkong.aliyuncs.com/hongfs/ecs-command-actions:main+g" .github/workflows/*
fi
