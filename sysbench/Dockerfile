FROM ghcr.io/hongfs/env:alpine as download

WORKDIR /code

RUN apk update && \
    apk add --no-cache wget && \
    wget https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.tar.gz -O /tmp/code.tar.gz && \
    tar -zxvf /tmp/code.tar.gz -C /tmp && \
    cd /tmp/sysbench-1.0.20 && \
    pwd

FROM ghcr.io/hongfs/env:ubuntu

WORKDIR /code

COPY --from=download /tmp/sysbench-1.0.20 /code

RUN apt-get update && \
    apt-get install -y make automake libtool pkg-config libaio-dev libmysqlclient-dev libssl-dev && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install
