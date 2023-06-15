FROM nginx:1.23.4

WORKDIR /web

COPY templates.conf /etc/nginx/conf.d/default.conf
COPY ratelimit.js /etc/nginx/ratelimit.js

RUN apt-get update && \
    apt-get install -y iproute2 curl tzdata ca-certificates && \
    sed -i '4i load_module modules/ngx_http_js_module.so;' /etc/nginx/nginx.conf
