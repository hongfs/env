FROM ghcr.io/hongfs/env:alpine-debug as base

WORKDIR /code

ADD https://github.com/percona/percona-toolkit/archive/refs/heads/3.x.tar.gz /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code --strip-components=1

FROM ghcr.io/hongfs/env:alpine-debug

WORKDIR /code

COPY --from=base /code /code

RUN apk update && \
    apk add --no-cache perl perl-dbd-mysql make && \
    rm -rf /var/cache/apk/* && \
    perl Makefile.PL && \
    make && \
    make install && \
    rm -rf /code/**

ENTRYPOINT ["pt-query-digest"]
