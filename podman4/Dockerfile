FROM quay.io/podman/stable:v4.7.0

RUN yum -y install buildah fuse fuse-overlayfs iputils nfs-utils libicu && \
    rm -rf /var/cache && \
    ln /usr/bin/podman /usr/bin/docker

