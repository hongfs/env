FROM ghcr.io/hongfs/env:ubuntu as download

WORKDIR /code/

RUN curl -s https://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/attach/113904/cn_zh/1580886874093/openapi-es-tools-master-b7a3ee97c5d23b5885d8e0ed5fa674b10f83d9e5.zip -o /tmp/edgescript.zip && \
    cd /tmp/ && \
    unzip /tmp/edgescript.zip && \
    ls -l && \
    cp -r /tmp/openapi-es-tools-master-b7a3ee97c5d23b5885d8e0ed5fa674b10f83d9e5/cdn-api/* /code/

FROM python:2.7.12-alpine

WORKDIR /code

COPY --from=download /code/ /code/

# docker run --rm -it --name edgescript ghcr.io/hongfs/env:aliyun-cdn-edgescript python es.py help
