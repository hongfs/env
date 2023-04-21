FROM ubuntu:22.04

WORKDIR /code

ADD https://awscli.amazonaws.com/awscli-exe-linux-x86_64-2.11.14.zip aws.zip

ENV AWS_EC2_METADATA_DISABLED=true

RUN apt-get update && \
    apt-get -y install unzip curl && \
    unzip aws.zip && \
    ./aws/install && \
    rm -rf aws.zip && \
    rm -rf /var/lib/apt/lists/* && \
    /usr/local/bin/aws configure set default.s3.addressing_style virtual && \
    /usr/local/bin/aws configure set default.s3.max_concurrent_requests 100
