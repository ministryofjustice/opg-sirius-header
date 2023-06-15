FROM golang:1.19.3-alpine as base
WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

FROM base as build-env

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o /go/bin/opg-sirius-header

FROM alpine:latest

WORKDIR /go/bin

RUN apk --update --no-cache add \
    ca-certificates && \
    apk upgrade busybox --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main && \
    apk upgrade libcrypto3 libssl3 && \
    rm -rf /var/cache/apk/*
RUN apk --no-cache add tzdata
ENV TZ="Europe/London"

COPY --from=build-env /go/bin/opg-sirius-header opg-sirius-header
COPY --from=build-env /app/static/template static/template
ENTRYPOINT ["./opg-sirius-header"]
