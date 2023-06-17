FROM alpine

WORKDIR /nginx

RUN apk update && \
    apk add --no-cache git curl gcc libc-dev make openssl-dev libpcre3-dev pcre2-dev zlib-dev linux-headers libxslt-dev gd-dev geoip-dev libedit-dev bash alpine-sdk findutils ca-certificates gettext && \
    cd && \
    git clone https://github.com/nginx/njs && \
    cd njs && \
    ./configure && \
    make njs && \
    cd && \
    git clone https://github.com/vozlt/nginx-module-vts && \
    git clone -b v0.3.2 https://github.com/vision5/ngx_devel_kit && \
    git clone -b v0.10.24 https://github.com/openresty/lua-nginx-module && \
    git clone -b v0.1.26 https://github.com/openresty/lua-resty-core && \
    git clone -b v0.13 https://github.com/openresty/lua-resty-lrucache && \
    git clone -b v2.1-20230410 https://github.com/openresty/luajit2 && \
    cd luajit2 && \
    make && \
    make install PREFIX=/usr/local/luajit2 && \
    export LUAJIT_LIB=/usr/local/luajit2/lib && \
    export LUAJIT_INC=/usr/local/luajit2/include/luajit-2.1 && \
    cd && \
    curl https://nginx.org/download/nginx-1.23.4.tar.gz -o nginx.tar.gz && \
    tar -zxvf nginx.tar.gz && \
    cd nginx-1.23.4 && \
    addgroup -g 101 -S nginx && \
    adduser -S -D -H -u 101 -h /var/cache/nginx -s /sbin/nologin -G nginx -g nginx nginx && \
    ./configure --prefix=/etc/nginx \
        --sbin-path=/usr/sbin/nginx \
        --modules-path=/usr/lib/nginx/modules \
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
        --with-ld-opt='-Wl,-rpath,/usr/local/luajit2/lib' \
        --add-module=$HOME/nginx-module-vts \
        --add-module=$HOME/ngx_devel_kit \
        --add-module=$HOME/lua-nginx-module \
        --add-module=$HOME/njs/nginx && \
    make -j2 && \
    make install && \
    mkdir -p /var/log/nginx && \
    mkdir -p /var/cache/nginx && \
    mkdir -p /var/cache/nginx/client_temp && \
    mkdir -p /var/cache/nginx/proxy_temp && \
    mkdir -p /var/cache/nginx/fastcgi_temp && \
    mkdir -p /var/cache/nginx/uwsgi_temp && \
    mkdir -p /var/cache/nginx/scgi_temp && \
    cd &&\
    cd lua-resty-core && \
    make install PREFIX=/etc/nginx && \
    cd &&\
    cd lua-resty-lrucache && \
    make install PREFIX=/etc/nginx && \
    cd &&\
    apk del git gcc make && \
    rm -rf /var/cache/apk && \
    rm -rf $HOME/nginx.tar.gz && \
    rm -rf $HOME/nginx-module-vts && \
    rm -rf $HOME/ngx_devel_kit && \
    rm -rf $HOME/lua-nginx-module && \
    rm -rf $HOME/lua-resty-core && \
    rm -rf $HOME/lua-resty-lrucache && \
    rm -rf $HOME/njs && \
    rm -rf $HOME/luajit2 && \
    rm -rf $HOME/nginx-1.23.4

COPY default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

STOPSIGNAL SIGQUIT

CMD ["nginx", "-g", "daemon off;"]
