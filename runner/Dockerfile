FROM ghcr.io/hongfs/env:ubuntu

WORKDIR /runner

RUN apt-get update && \
    apt-get install -y podman && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir actions-runner && \
    cd actions-runner && \
    curl -o runner.tar.gz -L https://github.com/actions/runner/releases/download/v2.310.2/actions-runner-linux-x64-2.310.2.tar.gz && \
    tar xzf ./runner.tar.gz && \
    rm ./runner.tar.gz && \
    pwd && \
    ls -lh
