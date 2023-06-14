FROM ghcr.io/hongfs/env:php82-fpm as base

WORKDIR /code

RUN composer create-project topthink/think tp

FROM ghcr.io/hongfs/env:php82-fpm

WORKDIR /code

COPY --from=base /code/tp /code
