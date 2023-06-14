FROM ghcr.io/hongfs/env:php82-cli as base

WORKDIR /code

RUN composer create-project topthink/think tp

FROM ghcr.io/hongfs/env:php82-cli

WORKDIR /code

COPY --from=base /code/tp /code

EXPOSE 8080

CMD ["php", "think", "run", "-p=8080"]
