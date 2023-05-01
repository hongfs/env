# 来源：https://github.com/skandyla/docker-wrk/blob/master/alpine/Dockerfile

FROM alpine as build

WORKDIR /code

RUN apk update && \
    apk --no-cache add alpine-sdk perl linux-headers && \
    git clone https://github.com/wg/wrk && \
    cd wrk && \
    make && \
    rm -rf /var/cache/apk/*

FROM alpine

COPY --from=build /code/wrk /usr/local/bin/

WORKDIR /data

ENTRYPOINT ["wrk"]