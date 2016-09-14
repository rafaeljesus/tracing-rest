FROM gliderlabs/alpine
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>

COPY . /go/src/github.com/rafaeljesus/tracing-rest

ADD db/dbconf.yml /db/dbconf.yml

RUN apk-install -t build-deps build-essential go git mercurial \
      && cd /go/src/github.com/rafaeljesus/tracing-rest \
      && export GOPATH=/go \
      && go get \
      && go get bitbucket.org/liamstask/goose/cmd/goose \
      && CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o /tracing-rest . \
      && rm -rf /go \
      && apk del --purge build-deps

ENTRYPOINT ["/tracing-rest"]
