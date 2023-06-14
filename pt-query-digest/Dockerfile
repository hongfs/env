FROM ghcr.io/hongfs/env:alpine-debug

WORKDIR /code

ADD https://github.com/percona/percona-toolkit/archive/refs/heads/3.x.tar.gz /tmp/

RUN apk update && \
    apk add perl perl-dbd-mysql && \
    apk add --virtual=build make && \
    cd /tmp && \
    ls -l
