# 编译
FROM ghcr.io/hongfs/env:golang120 as build

WORKDIR /code

COPY . .

RUN go mod tidy && \
    env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM ghcr.io/hongfs/env:alpine

WORKDIR /build

COPY --from=build /code/main ./

EXPOSE 9000

CMD ["./main"]