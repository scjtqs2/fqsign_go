FROM golang:1.18-alpine3.16 AS builder

#RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache git \
  && go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY ./ .

RUN set -ex \
    && BUILD=`date +%FT%T%z` \
    && COMMIT_SHA1=`git rev-parse HEAD` \
    && go build -ldflags "-s -w -extldflags '-static' -X main.Version=${COMMIT_SHA1}" -X main.Build=${BUILD} -v -o fqsign ./cmd/server/* \
    && mv config.yaml.dist config.yaml


FROM alpine:3.16 AS production

#RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY --from=builder /build/fqsign /app/fqsign
COPY --from=builder /build/config.yaml /app/config.yaml

RUN chmod +x /app/fqsign

CMD ["/app/fqsign","-c","/app/config.yaml"]