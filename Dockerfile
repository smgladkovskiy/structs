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

COPY . /go/src/gitlab.teamc.io/teamc.io/golang/nulls

WORKDIR /go/src/gitlab.teamc.io/teamc.io/golang/nulls

RUN make deps