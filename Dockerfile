FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD tracing-rest /tracing-rest
ENV TRACKING_REST_DB="dbname=tracing_rest_dev sslmode=disable"
ENV TRACKING_REST_PORT="3000"
ENTRYPOINT ["/tracing-rest"]

# FROM gliderlabs/alpine
# MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
#
# COPY . /go/src/github.com/rafaeljesus/tracing-rest
#
# ADD db/dbconf.yml /db/dbconf.yml
#
# RUN set -ex \
#       && apk add --no-cache --virtual .build-deps \
#       bash \
#       gcc \
#       musl-dev \
#       git \
#       openssl \
#       go \
#       \
#       && cd /go/src/github.com/rafaeljesus/tracing-rest \
#       && export GOPATH=/go \
#       && go get \
#       && CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o /tracing-rest . \
#       && rm -rf /go \
#       && apk del --purge build-deps
#
# ENTRYPOINT ["/tracing-rest"]
