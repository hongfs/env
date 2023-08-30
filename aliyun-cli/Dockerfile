FROM ubuntu:22.04 AS download

WORKDIR /code

ADD https://github.com/aliyun/aliyun-cli/releases/download/v3.0.160/aliyun-cli-linux-3.0.160-amd64.tgz /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code && \
    rm -rf /tmp/code.tar.gz

CMD [ "aliyun" ]