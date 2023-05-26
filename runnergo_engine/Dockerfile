FROM --platform=$TARGETPLATFORM ubuntu:23.04 AS github

WORKDIR /code

ADD https://github.com/Runner-Go-Team/RunnerGo-engine-open/archive/refs/heads/open.tar.gz /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code --strip-components=1

FROM --platform=$TARGETPLATFORM golang:1.20.4 AS build

WORKDIR /code

COPY --from=github /code/ /code/

RUN go mod tidy && \
    env GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main main.go

FROM --platform=$TARGETPLATFORM alpine:3.7 as run

WORKDIR /data

COPY --from=build /code/main .

EXPOSE 30000

CMD ["./main","-m", "1"]
