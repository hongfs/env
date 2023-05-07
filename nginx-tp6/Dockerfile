FROM nginx:1.23.4

WORKDIR /web

COPY templates.conf /etc/nginx/templates/default.conf

RUN apt-get update && \
    apt-get install -y iproute2 && \
    rm -rf /var/lib/apt/lists/*
