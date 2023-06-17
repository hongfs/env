FROM openresty/openresty:1.21.4.1-7-alpine-apk

COPY default.conf /etc/nginx/conf.d/default.conf

RUN apk update && \
    apk add git && \
    mkdir -p /etc/nginx/lua && \
    cd /etc/nginx/lua && \
    git clone https://github.com/timebug/lua-resty-redis-ratelimit && \
    git clone https://github.com/openresty/lua-resty-redis && \
    cd && \
    rm -rf /var/cache/apk
