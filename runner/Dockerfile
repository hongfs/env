FROM quay.io/podman/stable:v4.9.3

WORKDIR /runner

ENV RUNNER_ALLOW_RUNASROOT=1
ENV RUNNER_VERSION=2.326.0

COPY entrypoint.sh /hongfs/entrypoint.sh

RUN yum -y install buildah fuse fuse-overlayfs iputils nfs-utils libicu && \
    rm -rf /var/cache && \
    ln /usr/bin/podman /usr/bin/docker && \
    chmod +x /hongfs/entrypoint.sh

RUN echo '#!/bin/bash' > /etc/rc.local && \
    echo "mount -a -t nfs" >> /etc/rc.local && \
    chmod +x /etc/rc.local && \
    echo "options sunrpc tcp_slot_table_entries=128" >> /etc/modprobe.d/sunrpc.conf && \
    echo "options sunrpc tcp_max_slot_table_entries=128" >>  /etc/modprobe.d/sunrpc.conf && \
    mkdir -p /mnt/cache

RUN mkdir actions-runner && \
    cd actions-runner && \
    curl -o runner.tar.gz -L https://github.com/actions/runner/releases/download/v$RUNNER_VERSION/actions-runner-linux-x64-$RUNNER_VERSION.tar.gz && \
    tar xzf ./runner.tar.gz && \
    rm -f ./runner.tar.gz && \
    # rm -rf ./externals && \
    pwd

ENTRYPOINT [ "/hongfs/entrypoint.sh" ]
