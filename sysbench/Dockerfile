FROM ubuntu:23.04

RUN apt-get update && \
    apt-get install -y git curl wget make automake libtool pkg-config libaio-dev libmysqlclient-dev libssl-dev \
    git clone "https://github.com/akopytov/sysbench" && \
    cd sysbench && \
    git checkout 1.0.20 && \
    ./autogen.sh && \
    ./configure --prefix=/usr --mandir=/usr/share/man && \
    make && \
    make install \
    rm -rf /var/lib/apt/lists/*
