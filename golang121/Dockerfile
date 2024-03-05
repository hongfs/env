FROM docker.io/library/golang:1.21.7-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOOS=linux

RUN apk update && \
    apk add --no-cache build-base && \
    rm -rf /var/cache/apk/*
