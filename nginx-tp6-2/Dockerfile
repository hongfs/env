FROM docker.io/library/nginx:1.27.0

WORKDIR /web

COPY templates.conf /etc/nginx/templates/default.conf

RUN apt-get update && \
    apt-get install -y iproute2 && \
    rm -rf /var/lib/apt/lists/*
