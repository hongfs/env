FROM docker.io/library/nginx:1.27.0

RUN sed -i '3i load_module "/etc/nginx/modules/ngx_http_geoip_module.so";' /etc/nginx/nginx.conf && \
    sed -i '3i load_module "/etc/nginx/modules/ngx_http_image_filter_module.so";' /etc/nginx/nginx.conf && \
    sed -i '3i load_module "/etc/nginx/modules/ngx_http_js_module.so";' /etc/nginx/nginx.conf
