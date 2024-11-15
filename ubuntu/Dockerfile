FROM docker.io/library/ubuntu:24.04

ENV TZ=Asia/Shanghai

RUN apt-get update && \
    apt-get install -y curl wget ca-certificates tzdata unzip zip && \
    rm -rf /var/lib/apt/lists/* && \
    echo "1"
