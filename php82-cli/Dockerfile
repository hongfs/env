FROM library/composer:2.5.8 as composer

FROM php:8.2.8-cli as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN apt-get update && \
    apt-get install -y git zip libpng-dev libzip-dev && \
    rm -rf /var/lib/apt/lists/* && \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    pecl install redis && \
    docker-php-ext-install pdo pdo_mysql mysqli gd zip bcmath opcache && \
    docker-php-ext-enable mysqli gd redis

# 设置 opcache
RUN echo "[opcache]\n\
opcache.enable=1\n\
opcache.enable_cli=1\n\
opcache.memory_consumption=256\n\
opcache.interned_strings_buffer=128\n\
opcache.max_accelerated_files=1000000\n\
opcache.max_wasted_percentage=10\n\
opcache.max_file_size=0\n\
opcache.consistency_checks=0\n\
opcache.file_update_protection=0\n\
opcache.validate_timestamps=0\n\
opcache.enable_file_override=0" > /usr/local/etc/php/conf.d/opcache.ini

