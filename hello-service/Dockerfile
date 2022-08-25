#syntax=docker/dockerfile:1

## Build
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
RUN go get ./...

COPY *.go ./
RUN go build -o /hello

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /hello /hello

USER nonroot:nonroot

ENTRYPOINT ["/hello"]
