FROM openresty/openresty:1.21.4.1-7-alpine-apk

RUN apk update && \
    apk add git && \
    rm -rf /var/cache/apk && \
    mkdir -p /etc/nginx/lua && \
    cd /etc/nginx/lua && \
    git clone https://github.com/timebug/lua-resty-redis-ratelimit && \
    git clone https://github.com/openresty/lua-resty-redis && \
    git clone https://github.com/openresty/lua-resty-mysql && \
    git clone https://github.com/openresty/lua-resty-dns && \
    git clone https://github.com/openresty/lua-resty-websocket && \
    git clone https://github.com/openresty/lua-resty-lock && \
    git clone https://github.com/openresty/lua-resty-lrucache && \
    git clone https://github.com/openresty/lua-resty-string && \
    pwd
