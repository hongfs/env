FROM ghcr.io/hongfs/env:ubuntu AS download

WORKDIR /code

ENV VERSION=3.0.249

ADD https://github.com/aliyun/aliyun-cli/releases/download/v$VERSION/aliyun-cli-linux-$VERSION-amd64.tgz /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code && \
    rm -rf /tmp/code.tar.gz && \
    cp /code/aliyun /usr/local/bin/aliyun

CMD [ "aliyun" ]