FROM debian:stable-20230502-slim

RUN apt-get update && \
    apt-get install -y iproute2 curl tzdata ca-certificates && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    rm -rf /var/lib/apt/lists/*
