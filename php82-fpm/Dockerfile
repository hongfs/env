FROM library/composer:2.5.5 as composer

RUN composer config -g repo.packagist composer https://hk.hongfs.dev/composer/

FROM library/php:8.2.4-fpm as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN apt-get update && \
    apt-get install -y git zip libpng-dev libzip-dev && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    pecl install redis-5.3.7 && \
    docker-php-ext-install pdo pdo_mysql mysqli gd zip bcmath opcache && \
    docker-php-ext-enable mysqli gd redis

RUN echo "[opcache]\n\
opcache.enable=0\n\
opcache.enable_cli=0" > /usr/local/etc/php/conf.d/opcache.ini
