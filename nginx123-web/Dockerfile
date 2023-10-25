FROM nginx:1.25.3

WORKDIR /web

ADD https://raw.githubusercontent.com/hongfs/hongfs/762cd92804b828d31d28c65e7fceecb45f61a245/php/release.conf /etc/nginx/templates/default.conf

ENV WEB_ROOT=/web

EXPOSE 80
EXPOSE 443

CMD envsubst '${WEB_DOMAIN},${WEB_ROOT},${WEB_SSL_PORT},${WEB_SSL_PEM},${WEB_SSL_KEY},${WEB_FASTCGI}' < /etc/nginx/templates/default.conf > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'
