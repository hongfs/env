FROM ghcr.io/hongfs/env:alpine

WORKDIR /code

RUN apk update && \
    apk add --no-cache git \
    git clone https://github.com/akopytov/sysbench.git && \
    cd sysbench && \
    pwd
