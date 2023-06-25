FROM ghcr.io/hongfs/env:alpine

WORKDIR /code

RUN apk update && \
    apk add --no-cache wget && \
    wget https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.tar.gz -O /tmp/code.tar.gz && \
    tar -zxvf /tmp/code.tar.gz -C /tmp && \
    cd /tmp/sysbench-1.0.20 && \
    pwd
