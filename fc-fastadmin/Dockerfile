FROM ghcr.io/hongfs/env:php82-cli

LABEL org.opencontainers.image.source="https://github.com/hongfs/env/tree/main/fc-fastadmin"
LABEL org.opencontainers.image.description "FastAdmin 线上环境，请使用 >= 1.5.0 版本"

RUN sed -i "s+session.use_cookies = 0+session.use_cookies = 1+" "$PHP_INI_DIR/php.ini" && \
    sed -i "s+session.use_only_cookies = 0+session.use_only_cookies = 1+" "$PHP_INI_DIR/php.ini" && \
    # composer install -vvv --ignore-platform-req=composer-plugin-api
    pwd
