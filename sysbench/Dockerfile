FROM ubuntu:23.04

WORKDIR /code

RUN apt-get update && \
    apt-get install -y git \
    git clone "git@github.com:akopytov/sysbench.git" && \
    cd sysbench && \
    pwd
