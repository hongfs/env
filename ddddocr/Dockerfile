FROM ubuntu:22.04 AS download

WORKDIR /code

ADD https://api.github.com/repos/sml2h3/ocr_api_server/tarball/main /tmp/code.tar.gz

RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

RUN curl -Ls https://api.github.com/repos/sml2h3/ocr_api_server/tarball/main -o /tmp/code.tar.gz && \
    mkdir /code && \
    tar -xzf /tmp/code.tar.gz -C /code --strip-components=1 && \
    rm -rf /tmp/code.tar.gz

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
