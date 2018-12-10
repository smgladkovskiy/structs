# Докерфайл для тестов
FROM golang:alpine
MAINTAINER Sergey Gladkovskiy <smgladkovskiy@gmail.com>

ARG DEP_VERSION="0.4.1"
ARG SRC="/go/src/github.com/smgladkovskiy/structs"

RUN apk update \
 && apk add --no-cache \
    ca-certificates \
    curl \
    make \
 && curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /bin/dep \
 && chmod +x /bin/dep \
 && rm -rf /var/cache/apk/* \
 && rm -rf /tmp/*

WORKDIR ${SRC}
COPY . ${SRC}
RUN make all