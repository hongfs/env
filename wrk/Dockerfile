# 来源：https://github.com/skandyla/docker-wrk/blob/master/alpine/Dockerfile

FROM alpine as build

WORKDIR /code

RUN apk update && \
    apk --no-cache add alpine-sdk perl linux-headers && \
    git clone --depth 1 https://github.com/wg/wrk && \
    cd wrk && \
    make && \
    cp wrk /usr/local/bin && \
    rm -rf ./.git && \
    rm -rf /var/cache/apk/*

ENTRYPOINT ["wrk"]