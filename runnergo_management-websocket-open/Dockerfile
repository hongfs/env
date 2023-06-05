FROM --platform=$TARGETPLATFORM ubuntu:23.04 AS github

WORKDIR /code

ADD https://github.com/Runner-Go-Team/RunnerGo-management-websocket-open/archive/b4da4091f9924196baaeb2ea756499acf27aa505.tar.gz /tmp/code.tar.gz

RUN tar -xzf /tmp/code.tar.gz -C /code --strip-components=1

FROM --platform=$TARGETPLATFORM golang:1.20.4 AS build

WORKDIR /code

COPY --from=github /code/ /code/

RUN go mod tidy && \
    env GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main main.go

FROM --platform=$TARGETPLATFORM alpine:3.7 as run

WORKDIR /data

COPY --from=build /code/main .

COPY ./configs ./configs

CMD ["./main","-m", "1"]
