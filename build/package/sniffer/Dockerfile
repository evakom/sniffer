FROM golang:1.13.8-alpine as builder
LABEL maintainer="Eugene Klimov"

RUN apk add --no-cache --virtual .build-deps alpine-sdk gcc libpcap-dev

WORKDIR /root
COPY . .
RUN go build -o app/sniffer cmd/sniffer/*.go

