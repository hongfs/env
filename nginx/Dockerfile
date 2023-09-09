FROM nginx:1.25.2-alpine

RUN sed -i '3i load_module "/etc/nginx/modules/ngx_http_geoip_module.so";' /etc/nginx/nginx.conf && \
    sed -i '3i load_module "/etc/nginx/modules/ngx_http_image_filter_module.so";' /etc/nginx/nginx.conf && \
    sed -i '3i load_module "/etc/nginx/modules/ngx_http_js_module.so";' /etc/nginx/nginx.conf
