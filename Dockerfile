# Докерфайл для тестов
FROM golang:alpine
MAINTAINER Sergey Gladkovskiy <smgladkovskiy@gmail.com>

RUN apk update \
 && apk add --no-cache \
    ca-certificates \
    curl \
    git \
    make \
    openssl \
 && rm -rf /var/cache/apk/* \
 && rm -rf /tmp/*

COPY . /go/src/github.com/smgladkovskiy/structs

WORKDIR /go/src/github.com/smgladkovskiy/structs

RUN make deps