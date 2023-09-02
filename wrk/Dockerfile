# 来源：https://github.com/skandyla/docker-wrk/blob/master/alpine/Dockerfile

FROM alpine as build

WORKDIR /code

RUN apk update && \
    apk --no-cache add alpine-sdk perl libssl-dev luajit linux-headers make gcc && \
    git clone https://github.com/wg/wrk && \
    cd wrk && \
    make && \
    cp wrk /usr/local/bin && \
    rm -rf /var/cache/apk/*

ENTRYPOINT ["wrk"]