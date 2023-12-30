FROM golang:1.21.5 as build

WORKDIR /code

COPY main.go main.go

RUN env GOOS=linux GOARCH=amd64 go build -o main main.go

# FROM ghcr.io/hongfs/env:alpine

# WORKDIR /build

# COPY --from=build /code/main /build/

EXPOSE 80

CMD ["./main"]