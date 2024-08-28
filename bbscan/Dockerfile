FROM ghcr.io/hongfs/env:python312

RUN apt-get update && \
    apt-get install -y git && \
    git clone https://github.com/lijiejie/BBScan && \
    cd BBScan && \
    pip3 install -r requirements.txt && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

CMD ["python", "BBScan.py"]
