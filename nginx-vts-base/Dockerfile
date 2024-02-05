FROM ghcr.io/hongfs/env:alpine

ENV NGINX_VERSION=1.25.3

RUN apk update && \
    apk add --no-cache git curl gcc libc-dev make openssl-dev pcre2-dev zlib-dev linux-headers libxslt-dev gd-dev geoip-dev libedit-dev bash alpine-sdk findutils ca-certificates gettext && \
    cd && \
    git clone https://github.com/nginx/njs && \
    cd njs && \
    ./configure && \
    make njs && \
    cd ../ && \
    git clone https://github.com/vozlt/nginx-module-vts && \
    curl https://nginx.org/download/nginx-$NGINX_VERSION.tar.gz -o nginx.tar.gz && \
    tar -zxvf nginx.tar.gz && \
    cd nginx-$NGINX_VERSION && \
    addgroup -g 101 -S nginx && \
    adduser -S -D -H -u 101 -h /var/cache/nginx -s /sbin/nologin -G nginx -g nginx nginx && \
    ./configure --prefix=/etc/nginx \
        --sbin-path=/usr/sbin/nginx \
        --modules-path=/etc/nginx/modules \
        --conf-path=/etc/nginx/nginx.conf \
        --error-log-path=/var/log/nginx/error.log \
        --http-log-path=/var/log/nginx/access.log \
        --pid-path=/var/run/nginx.pid \
        --lock-path=/var/run/nginx.lock \
        --http-client-body-temp-path=/var/cache/nginx/client_temp \
        --http-proxy-temp-path=/var/cache/nginx/proxy_temp \
        --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp \
        --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp \
        --http-scgi-temp-path=/var/cache/nginx/scgi_temp \
        --user=nginx \
        --group=nginx \
        --with-compat \
        --with-file-aio \
        --with-threads \
        --with-http_addition_module \
        --with-http_auth_request_module \
        --with-http_dav_module \
        --with-http_flv_module \
        --with-http_gunzip_module \
        --with-http_gzip_static_module \
        --with-http_random_index_module \
        --with-http_realip_module \
        --with-http_secure_link_module \
        --with-http_slice_module \
        --with-http_ssl_module \
        --with-http_stub_status_module \
        --with-http_sub_module \
        --with-http_v2_module \
        --with-cc-opt='-g -O2 -fstack-protector-strong -Wformat -Werror=format-security -Wp,-D_FORTIFY_SOURCE=2 -fPIC' \
        --with-ld-opt='-Wl,-z,relro -Wl,-z,now -Wl,--as-needed -pie' \
        --add-module=$HOME/nginx-module-vts \
        --add-module=$HOME/njs/nginx && \
    make && \
    make install && \
    make modules && \
    mkdir -p /var/log/nginx && \
    mkdir -p /var/cache/nginx && \
    mkdir -p /var/cache/nginx/client_temp && \
    mkdir -p /var/cache/nginx/proxy_temp && \
    mkdir -p /var/cache/nginx/fastcgi_temp && \
    mkdir -p /var/cache/nginx/uwsgi_temp && \
    mkdir -p /var/cache/nginx/scgi_temp && \
    apk del git gcc make && \
    rm -rf /var/cache/apk && \
    rm -rf $HOME/nginx.tar.gz && \
    rm -rf $HOME/nginx-module-vts && \
    rm -rf $HOME/njs && \
    rm -rf $HOME/nginx-$NGINX_VERSION

EXPOSE 80

STOPSIGNAL SIGQUIT

CMD ["nginx", "-g", "daemon off;"]
