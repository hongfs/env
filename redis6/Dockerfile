FROM docker.io/library/redis:6.2.14

RUN apt-get update && \
    apt-get install -y iproute2 curl tzdata ca-certificates iputils-ping coreutils && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    rm -rf /var/lib/apt/lists/*

