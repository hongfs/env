FROM ghcr.io/hongfs/env:php82-cli as php

# https://github.com/Unitech/pm2
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash && \
    nvm install 20.1.0 && \
    npm config set prefix /usr/local && \
    npm install -g npm && \
    npm install -g pm2 && \
    npm version
