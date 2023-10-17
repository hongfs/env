FROM quay.io/podman/stable:v4.7.0

WORKDIR /runner

ENV RUNNER_ALLOW_RUNASROOT=1

RUN yum -y install podman buildah fuse fuse-overlayfs nfs-utils libicu && \
    rm -rf /var/cache/yum

RUN echo '#!/bin/bash' > /etc/rc.local && \
    echo "mount -a -t nfs" >> /etc/rc.local && \
    chmod +x /etc/rc.local && \
    echo "options sunrpc tcp_slot_table_entries=128" >> /etc/modprobe.d/sunrpc.conf && \
    echo "options sunrpc tcp_max_slot_table_entries=128" >>  /etc/modprobe.d/sunrpc.conf && \
    mkdir -P /mnt/cache

RUN mkdir actions-runner && \
    cd actions-runner && \
    curl -o runner.tar.gz -L https://github.com/actions/runner/releases/download/v2.310.2/actions-runner-linux-x64-2.310.2.tar.gz && \
    tar xzf ./runner.tar.gz && \
    rm ./runner.tar.gz && \
    # rm -rf ./externals && \
    pwd
