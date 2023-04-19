FROM ubuntu:22.04

LABEL org.opencontainers.image.description aws-cli-20230419

WORKDIR /code

ADD https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip aws.zip

RUN apt-get update && \
    apt-get -y install unzip curl && \
    unzip aws.zip && \
    ./aws/install && \
    rm -rf aws.zip && \
    rm -rf /var/lib/apt/lists/* && \
    echo "1"
