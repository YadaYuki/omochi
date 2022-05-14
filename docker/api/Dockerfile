FROM golang:1.18-alpine3.14 AS build 

ENV GO111MODULE=on

WORKDIR /go/github.com/YadaYuki/omochi

COPY ./go.mod ./go.sum ./

RUN mkdir shell
COPY ./shell ./shell


RUN apk update \
    && apk add mariadb-client

RUN go mod download

COPY ./ ./
