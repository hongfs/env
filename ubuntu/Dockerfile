FROM ubuntu:22.04

ENV TZ=Asia/Shanghai

RUN apt-get update && \
    apt-get install -y curl wget ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/*
