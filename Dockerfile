FROM golang:1.14.2-alpine3.11

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY ./.bin/app .
COPY ./config/ ./config/