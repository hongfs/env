FROM library/composer:2.5.5 as composer

FROM library/php:8.2.5-fpm as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN apt-get update && \
    apt-get install -y git zip libpng-dev libzip-dev libmagickwand-dev iproute2 sysctl && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    pecl install redis-5.3.7 && \
    pecl install imagick-3.7.0 && \
    docker-php-ext-install pdo pdo_mysql mysqli gd zip bcmath opcache && \
    docker-php-ext-enable mysqli gd redis imagick

# 设置 opcache
RUN echo "[opcache]\n\
opcache.enable=1\n\
opcache.enable_cli=1\n\
opcache.memory_consumption=256\n\
opcache.interned_strings_buffer=128\n\
opcache.max_accelerated_files=1000000\n\
opcache.max_wasted_percentage=10\n\
opcache.validate_timestamps=0\n\
opcache.enable_file_override=1" > /usr/local/etc/php/conf.d/opcache.ini

# 设置 php-fpm
RUN echo "\n\
pm = static\n\
pm.max_children = 20\n\
pm.max_requests = 50000" >> /usr/local/etc/php-fpm.d/zz-docker.conf
