FROM golang:1.17.8-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go build -o /bye 

FROM alpine:3.16.2

WORKDIR /
COPY --from=build /bye /bye

ENTRYPOINT ["./bye"]
