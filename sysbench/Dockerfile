FROM ghcr.io/hongfs/env:ubuntu AS download

WORKDIR /code

ADD https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.tar.gz /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code --strip-components=1

FROM ghcr.io/hongfs/env:ubuntu

WORKDIR /code

COPY --from=download /code /code

RUN apt-get update && \
    apt-get install -y make automake libtool pkg-config libaio-dev libmysqlclient-dev libssl-dev coreutils \
    ls && \
    ./autogen.sh && \
    ./configure --prefix=/usr --mandir=/usr/share/man && \
    make && \
    make install \
    rm -rf /var/lib/apt/lists/*
