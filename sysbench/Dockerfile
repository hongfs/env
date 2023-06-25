FROM ghcr.io/hongfs/env:alpine

WORKDIR /code

RUN apk update && \
    apk add --no-cache github-cli \
    gh repo clone akopytov/sysbench && \
    cd sysbench && \
    pwd
