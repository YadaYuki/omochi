FROM golang:1.18-alpine3.14 AS build 

ENV GO111MODULE=on

WORKDIR /go/github.com/YadaYuki/omochi/app

COPY ./app/go.mod ./app/go.sum ./

RUN apk update \
    && apk add mariadb-client \
    && apk add bash

RUN go mod download

