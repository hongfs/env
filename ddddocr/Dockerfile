FROM ubuntu:22.04 AS download

WORKDIR /code

ADD https://github.com//archive/refs/heads/main.tgz https://api.github.com/repos/sml2h3/ocr_api_server/tarball/main /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code && \
    rm -rf /tmp/code.tar.gz && \
    pwd

FROM python:3.8-slim-buster

RUN mkdir /app

COPY --from=download /code/*.txt /code/*.py /code/*.sh /code/*.onnx /app/

RUN cd /app && \
    python3 -m pip install --upgrade pip && \
    pip3 install --no-cache-dir -r requirements.txt && \
    rm -rf /tmp/* && rm -rf /root/.cache/* && \
    apt-get --allow-releaseinfo-change update && apt install libgl1-mesa-glx libglib2.0-0 -y && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

EXPOSE 9898

CMD ["python3", "ocr_server.py", "--port", "9898", "--ocr", "--det"]
