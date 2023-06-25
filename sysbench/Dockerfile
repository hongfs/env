FROM ubuntu:23.04

WORKDIR /code

RUN apt-get update && \
    apt-get install -y tar curl wget make automake libtool pkg-config libaio-dev libmysqlclient-dev libssl-dev coreutils \
    curl https//github.com/akopytov/sysbench/archive/refs/tags/1.0.20.tar.gz -o /tmp/code.tar.gz && \
    tar -xzf /tmp/code.tar.gz -C /code --strip-components=1 && \
    cd /tmp/ && \
    ls && \
    ./autogen.sh && \
    ./configure --prefix=/usr --mandir=/usr/share/man && \
    make && \
    make install \
    rm -rf /var/lib/apt/lists/*
