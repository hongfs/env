FROM docker.io/library/ubuntu:24.04

WORKDIR /build

ENV K6_VERSION=0.53.0

ADD https://github.com/grafana/k6/releases/download/v$K6_VERSION/k6-v$K6_VERSION-linux-amd64.tar.gz ./

RUN tar -xvf k6-v$K6_VERSION-linux-amd64.tar.gz && \
    mv k6-v$K6_VERSION-linux-amd64/k6 /usr/local/bin && \
    rm -rf k6-v$K6_VERSION-linux-amd64 && \
    rm k6-v$K6_VERSION-linux-amd64.tar.gz

CMD ["/usr/local/bin/k6"]
