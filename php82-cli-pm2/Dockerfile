FROM ghcr.io/hongfs/env:php82-cli as php

ENV NODE_VERSION 20.1.0
ENV NVM_NODEJS_ORG_MIRROR=https://repo.huaweicloud.com/nodejs/

# https://github.com/Unitech/pm2
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash && \
    export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")" && \
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" && \
    nvm install $NODE_VERSION && \
    nvm alias default $NODE_VERSION && \
    nvm use default && \
    npm install -g npm && \
    npm install -g pm2 && \
    npm version && \
    npm config ls -l && \
    ln $(npm config get prefix)/bin/node /usr/local/bin/node && \
    ln $(npm config get prefix)/bin/npm /usr/local/bin/npm && \
    ln $(npm config get prefix)/bin/pm2 /usr/local/bin/pm2 && \
    ln $(npm config get prefix)/bin/pm2-dev /usr/local/bin/pm2-dev && \
    ln $(npm config get prefix)/bin/pm2-docker /usr/local/bin/pm2-docker && \
    ln $(npm config get prefix)/bin/pm2-runtime /usr/local/bin/pm2-runtime
