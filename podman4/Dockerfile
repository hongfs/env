FROM quay.io/podman/stable:v4.7.0

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
