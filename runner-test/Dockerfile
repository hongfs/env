FROM ghcr.io/hongfs/env:podman4

WORKDIR /runner

ENV RUNNER_ALLOW_RUNASROOT=1
ENV RUNNER_VERSION=2.312.0

COPY entrypoint.sh /hongfs/entrypoint.sh

COPY update_ghcr_image.sh /hongfs/update_ghcr_image.sh

RUN chmod +x /hongfs/entrypoint.sh && \
    chmod +x /hongfs/update_ghcr_image.sh && \
    mkdir actions-runner && \
    cd actions-runner && \
    curl -o runner.tar.gz -L https://github.com/actions/runner/releases/download/v$RUNNER_VERSION/actions-runner-linux-x64-$RUNNER_VERSION.tar.gz && \
    tar xzf ./runner.tar.gz && \
    rm -f ./runner.tar.gz && \
    pwd

ENTRYPOINT [ "/hongfs/entrypoint.sh" ]
