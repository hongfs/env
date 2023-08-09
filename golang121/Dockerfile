FROM golang:1.21.0-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=linux

RUN apk update && \
    apk add --no-cache build-base && \
    rm -rf /var/cache/apk/*