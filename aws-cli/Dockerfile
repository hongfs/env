FROM ghcr.io/hongfs/env:debian

ENV AWS_EC2_METADATA_DISABLED=true

ENV VERSION=2.25.11

RUN apt-get update && \
    apt-get -y install unzip curl && \
    curl -Ls https://awscli.amazonaws.com/awscli-exe-linux-x86_64-$VERSION.zip -o aws.zip && \
    unzip aws.zip && \
    rm -rf aws.zip && \
    ./aws/install && \
    rm -rf ./aws && \
    rm -rf /var/lib/apt/lists/* && \
    /usr/local/bin/aws configure set default.s3.addressing_style virtual && \
    /usr/local/bin/aws configure set default.s3.max_concurrent_requests 100
