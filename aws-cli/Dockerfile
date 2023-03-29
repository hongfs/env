FROM ubuntu:22.04

WORKDIR /code

ADD https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip aws.zip

RUN unzip aws.zip && \
    ./aws/install && \
    rm -rf aws.zip
