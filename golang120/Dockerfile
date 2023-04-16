FROM golang:1.20.3-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1

RUN apk update && \
    apk add --no-cache build-base && \
    rm -rf /var/cache/apk/*