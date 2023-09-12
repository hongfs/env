FROM library/composer:2.5.8 as composer

FROM php:8.2.10-cli as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN apt-get update && \
    apt-get install -y git zip libevent-dev libpng-dev libzip-dev libssl-dev libmagickwand-dev iproute2 procps && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN docker-php-ext-install -j$(nproc) sockets && \
    docker-php-ext-install -j$(nproc) pdo && \
    docker-php-ext-install -j$(nproc) pdo_mysql && \
    docker-php-ext-install -j$(nproc) mysqli && \
    docker-php-ext-install -j$(nproc) gd && \
    docker-php-ext-install -j$(nproc) zip && \
    docker-php-ext-install -j$(nproc) bcmath && \
    docker-php-ext-install -j$(nproc) opcache && \
    docker-php-ext-install -j$(nproc) pcntl && \
    pecl install redis-6.0.0 && \
    pecl install event-3.0.8 && \
    pecl install msgpack-2.2.0 && \
    pecl install yac-2.3.1 && \
    pecl install yaconf-1.1.2 && \
    docker-php-ext-enable mysqli gd redis event msgpack yac yaconf

RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    # 关闭 PHP 版本的显示
    sed -i "s+expose_php = On+expose_php = Off+" "$PHP_INI_DIR/php.ini" && \
    # 序列化浮点数时存储的有效数字的位数。
    sed -i "s+serialize_precision = -1+serialize_precision = 14+" "$PHP_INI_DIR/php.ini" && \
    # 默认上海时区
    sed -i "s+;date.timezone =+date.timezone =Asia/Shanghai+" "$PHP_INI_DIR/php.ini" && \
    # BC 的小数点默认 2 位
    sed -i "s+bcmath.scale = 0+bcmath.scale = 2+" "$PHP_INI_DIR/php.ini" && \
    # 关闭 cookice
    sed -i "s+session.use_cookies = 1+session.use_cookies = 0+" "$PHP_INI_DIR/php.ini" && \
    sed -i "s+session.use_only_cookies = 1+session.use_only_cookies = 0+" "$PHP_INI_DIR/php.ini"

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
opcache.enable_file_override=0" > "$PHP_INI_DIR/conf.d/opcache.ini"

RUN php -m
