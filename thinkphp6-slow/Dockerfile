FROM ghcr.io/hongfs/env:php82-cli as base

WORKDIR /code

RUN composer create-project topthink/think tp

FROM ghcr.io/hongfs/env:php82-cli

WORKDIR /code

COPY --from=base /code/tp /code
COPY app.php config/app.php
COPY database.php config/database.php
COPY Index.php app/controller/Index.php

EXPOSE 8000

CMD ["sh", "-c", "php think run"]
