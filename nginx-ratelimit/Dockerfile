FROM debian:stable-20230612 as base

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

FROM nginx:1.23.4

WORKDIR /web

COPY templates.conf /etc/nginx/conf.d/default.conf
COPY ratelimit.js /etc/nginx/ratelimit.js

RUN apt-get update && \
    apt-get install -y iproute2 curl tzdata ca-certificates openssl && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    rm -rf /var/lib/apt/lists/* && \
    sed -i '4i load_module modules/ngx_http_js_module.so;' /etc/nginx/nginx.conf

COPY --from=base /etc/ca-certificates /etc/ca-certificates
