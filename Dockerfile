# Докерфайл для тестов
FROM golang:alpine as build-env
MAINTAINER Sergey Gladkovskiy <smgladkovskiy@gmail.com>

ARG DEP_VERSION="0.4.1"

RUN apk update \
 && apk add --no-cache \
    ca-certificates \
    curl \
    git \
    make \
    openssl \
 && curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /bin/dep \
 && chmod +x /bin/dep \
 && rm -rf /var/cache/apk/* \
 && rm -rf /tmp/*

COPY . /go/src/gitlab.teamc.io/teamc.io/golang/nulls

WORKDIR /go/src/gitlab.teamc.io/teamc.io/golang/nulls

RUN make deps