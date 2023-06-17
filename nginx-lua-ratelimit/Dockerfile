FROM openresty/openresty:1.21.4.1-0-jammy

WORKDIR /web

COPY default.conf /etc/nginx/conf.d/default.conf

RUN apt-get update && \
    apt-get install -y iproute2 curl tzdata ca-certificates openssl git && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    mkdir -p /etc/nginx/lua/ && \
    cd /etc/nginx/lua/ && \
    git clone https://github.com/timebug/lua-resty-redis-ratelimit && \
    git clone https://github.com/openresty/lua-resty-redis && \
    git clone https://github.com/openresty/lua-resty-limit-traffic && \
    rm -rf /var/lib/apt/lists/*
