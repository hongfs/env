FROM ghcr.io/hongfs/env:ubuntu

RUN apt-get update && \
    apt-get install -y git curl wget make automake libtool pkg-config libaio-dev libmysqlclient-dev libssl-dev \
    git clone -b 1.0.20 https://github.com/akopytov/sysbench && \
    cd sysbench && \
    ./autogen.sh && \
    ./configure --prefix=/usr --mandir=/usr/share/man && \
    make && \
    make install
