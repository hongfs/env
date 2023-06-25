FROM ubuntu:23.04

WORKDIR /code

RUN apt-get update && \
    apt-get install -y git tar curl wget make automake libtool pkg-config libaio-dev libmysqlclient-dev libssl-dev coreutils \
    git clone git@github.com:akopytov/sysbench.git && \
    cd sysbench && \
    pwd
