FROM openresty/openresty:1.21.4.1-7-alpine-apk

COPY default.conf /etc/nginx/conf.d/default.conf

RUN apk update && \
    apk add git && \
    mkdir -p /etc/nginx/lua && \
    cd && \
    git clone https://github.com/timebug/lua-resty-redis-ratelimit && \
    git clone https://github.com/openresty/lua-resty-redis && \
    git clone https://github.com/openresty/lua-resty-limit-traffic && \
    cd && \
    rm -rf /var/cache/apk
