# 编译
FROM ghcr.io/hongfs/env:golang121 as build

WORKDIR /code

COPY . .

RUN go mod tidy && \
    go build -o main main.go

FROM ghcr.io/hongfs/env:alpine

WORKDIR /build

COPY --from=build /code/main /build/

CMD ["./main"]