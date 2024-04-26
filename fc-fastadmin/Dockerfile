FROM ghcr.io/hongfs/env:php82-cli

RUN sed -i "s+session.use_cookies = 0+session.use_cookies = 1+" "$PHP_INI_DIR/php.ini" && \
    sed -i "s+session.use_only_cookies = 0+session.use_only_cookies = 1+" "$PHP_INI_DIR/php.ini" && \
    # composer install -vvv --ignore-platform-req=composer-plugin-api
    pwd
