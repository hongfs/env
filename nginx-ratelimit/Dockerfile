FROM nginx:1.23.4

WORKDIR /web

COPY templates.conf /etc/nginx/templates/default.conf
COPY ratelimit.js /etc/nginx/ratelimit.js

RUN sed -i '4i load_module modules/ngx_http_js_module.so;' /etc/nginx/nginx.conf
