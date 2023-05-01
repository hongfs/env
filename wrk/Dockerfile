# 来源：https://github.com/skandyla/docker-wrk/blob/master/alpine/Dockerfile

FROM ghcr.io/hongfs/env:alpine as build

WORKDIR /code

RUN apk add --update alpine-sdk perl linux-headers

RUN git clone https://github.com/wg/wrk && \
    cd wrk && \
    make

FROM ghcr.io/hongfs/env:alpine

COPY --from=build /code/wrk /usr/local/bin/

RUN apk add --no-cache libgcc

WORKDIR /data

ENTRYPOINT ["wrk"]