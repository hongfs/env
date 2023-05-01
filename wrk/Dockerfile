# 来源：https://github.com/skandyla/docker-wrk/blob/master/alpine/Dockerfile

FROM ghcr.io/hongfs/env:alpine as build

WORKDIR /code

RUN apk update && \
    apk add alpine-sdk perl linux-headers && \
    git clone https://github.com/wg/wrk && \
    cd wrk && \
    make && \
    rm -rf /var/cache/apk/*

FROM ghcr.io/hongfs/env:alpine

COPY --from=build /code/wrk /usr/local/bin/

RUN apk update && \
    apk add libgcc && \
    rm -rf /var/cache/apk/*

WORKDIR /data

ENTRYPOINT ["wrk"]