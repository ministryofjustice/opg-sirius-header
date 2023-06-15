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

FROM node:19.1.0-alpine3.16 as asset-env

WORKDIR /app

COPY static static
COPY package.json .
COPY yarn.lock .

RUN yarn install
RUN yarn build

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
COPY --from=build-env /app/web/template web/template
COPY --from=asset-env /app/static static

ENTRYPOINT ["./opg-sirius-header"]
