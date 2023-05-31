FROM library/composer:2.5.7 as composer

FROM library/php:8.2.5-fpm as php

COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN pwd && \
    mkdir -p /tmp/opcache && \
    apt-get update && \
    apt-get install -y git zip libpng-dev libzip-dev libmagickwand-dev iproute2 procps && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" && \
    # 关闭 PHP 版本的显示
    sed -i "s+expose_php = On+expose_php = Off+" "$PHP_INI_DIR/php.ini" && \
    # 关闭 GC，如果涉及到队列那可能会存在问题
    sed -i "s+zend.enable_gc = On+zend.enable_gc = Off+" "$PHP_INI_DIR/php.ini" && \
    # 序列化浮点数时存储的有效数字的位数。
    sed -i "s+serialize_precision = -1+serialize_precision = 14+" "$PHP_INI_DIR/php.ini" && \
    # 默认上海时区
    sed -i "s+;date.timezone =+date.timezone =Asia/Shanghai+" "$PHP_INI_DIR/php.ini" && \
    # BC 的小数点默认 2 位
    sed -i "s+bcmath.scale = 0+bcmath.scale = 2+" "$PHP_INI_DIR/php.ini" && \
    # 关闭 cookice
    sed -i "s+session.use_cookies = 1+session.use_cookies = 0+" "$PHP_INI_DIR/php.ini" && \
    sed -i "s+session.use_only_cookies = 1+session.use_only_cookies = 0+" "$PHP_INI_DIR/php.ini" && \
    pecl install redis-5.3.7 && \
    pecl install imagick-3.7.0 && \
    pecl install yaconf-1.1.2 && \
    pecl install yac-2.3.1 && \
    docker-php-ext-install pdo pdo_mysql mysqli gd zip bcmath opcache && \
    docker-php-ext-enable mysqli gd redis imagick yaconf yac

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

# 设置 php-fpm
RUN echo "\n\
pm = static\n\
pm.max_children = 20\n\
pm.max_requests = 10000" >> /usr/local/etc/php-fpm.d/zz-docker.conf
