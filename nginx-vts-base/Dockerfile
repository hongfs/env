FROM alpine

RUN apk update && \
    apk add --no-cache --virtual .checksum-deps git curl gcc libc-dev make openssl-dev pcre2-dev zlib-dev linux-headers libxslt-dev gd-dev geoip-dev libedit-dev bash alpine-sdk findutils ca-certificates && \
    cd && \
    git clone https://github.com/vozlt/nginx-module-vts.git && \
    curl https://repo.huaweicloud.com/nginx/nginx-1.23.4.tar.gz -o nginx.tar.gz && \
    tar -zxvf nginx.tar.gz && \
    cd nginx-1.23.4 && \
    ./configure --prefix=/etc/nginx --sbin-path=/usr/sbin/nginx --modules-path=/usr/lib/nginx/modules --conf-path=/etc/nginx/nginx.conf --error-log-path=/var/log/nginx/error.log --http-log-path=/var/log/nginx/access.log --pid-path=/var/run/nginx.pid --lock-path=/var/run/nginx.lock --http-client-body-temp-path=/var/cache/nginx/client_temp --http-proxy-temp-path=/var/cache/nginx/proxy_temp --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp --http-scgi-temp-path=/var/cache/nginx/scgi_temp --user=nginx --group=nginx --with-compat --with-file-aio --with-threads --with-http_addition_module --with-http_auth_request_module --with-http_dav_module --with-http_flv_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_random_index_module --with-http_realip_module --with-http_secure_link_module --with-http_slice_module --with-http_ssl_module --with-http_stub_status_module --with-http_sub_module --with-http_v2_module --add-module=/root/nginx-module-vts && \
    make && \
    make install && \
    apk del .checksum-deps

CMD [nginx, '-g', 'daemon off;']
