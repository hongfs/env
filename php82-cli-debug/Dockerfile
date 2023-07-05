FROM library/composer:2.5.8 as composer

FROM php:8.2.5-cli as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN apt-get update && \
    apt-get install -y git zip libpng-dev libzip-dev && \
    rm -rf /var/lib/apt/lists/* && \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    pecl install redis xhprof && \
    pecl install https://pecl.php.net/get/xdebug-3.2.1.tgz && \
    docker-php-ext-install pdo pdo_mysql mysqli gd zip bcmath && \
    docker-php-ext-enable mysqli gd redis xhprof xdebug

RUN git clone -b v2.3.9 https://github.com/longxinH/xhprof.git /xhprof_code

RUN echo "[xhprof]\n\
\n\
extension=xhprof.so\n\
xhprof.collect_additional_info=1\n\
xhprof.output_dir=/xhprof" >> "$PHP_INI_DIR/php.ini"

RUN echo "zend_extension=xdebug\n\
\n\
[Xdebug]\n\
xdebug.mode=debug,develop,trace\n\
xdebug.use_compression=false\n\
xdebug.trace_format=1\n\
xdebug.trace_output_name=%H_%u\n\
xdebug.remote_enable=true\n\
xdebug.client_port=9000\n\
xdebug.client_host=host.docker.internal\n\
xdebug.output_dir=/xdebug\n\
xdebug.collect_assignments=true\n\
xdebug.collect_return=true" > /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini
