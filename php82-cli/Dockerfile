FROM library/composer:2.5.8 as composer

FROM php:8.2.8-cli as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN apt-get update && \
    apt-get install -y git zip libpng-dev libzip-dev && \
    rm -rf /var/lib/apt/lists/* && \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    pecl install redis && \
    docker-php-ext-install pdo pdo_mysql mysqli gd zip bcmath && \
    docker-php-ext-enable mysqli gd redis

