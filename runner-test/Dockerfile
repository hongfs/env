FROM quay.io/podman/stable:v4.7.0

WORKDIR /runner

ENV RUNNER_ALLOW_RUNASROOT=1
ENV RUNNER_VERSION=2.312.0

COPY entrypoint.sh /hongfs/entrypoint.sh

COPY update_ghcr_image.sh /hongfs/update_ghcr_image.sh

ENV ACTIONS_RUNNER_HOOK_JOB_STARTED=/hongfs/update_ghcr_image.sh

RUN yum -y install buildah fuse fuse-overlayfs iputils nfs-utils libicu && \
    rm -rf /var/cache && \
    ln /usr/bin/podman /usr/bin/docker && \
    chmod +x /hongfs/entrypoint.sh && \
    chmod +x /hongfs/update_ghcr_image.sh && \
    mkdir actions-runner && \
    cd actions-runner && \
    curl -o runner.tar.gz -L https://github.com/actions/runner/releases/download/v$RUNNER_VERSION/actions-runner-linux-x64-$RUNNER_VERSION.tar.gz && \
    tar xzf ./runner.tar.gz && \
    rm -f ./runner.tar.gz && \
    pwd

ENTRYPOINT [ "/hongfs/entrypoint.sh" ]
