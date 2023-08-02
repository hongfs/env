FROM ghcr.io/hongfs/env:php82-cli as php

ENV NODE_VERSION 20.1.0

# https://github.com/Unitech/pm2
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | NVM_DIR="/usr/local/nvm" bash && \
    nvm install $NODE_VERSION && \
    nvm alias default $NODE_VERSION && \
    nvm use default && \
    npm install -g npm && \
    npm install -g pm2 && \
    npm version && \
    npm config ls -l
