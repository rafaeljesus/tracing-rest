FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD tracing-rest /tracing-rest
ENV TRACKING_REST_DB="postgres://postgres:@docker/tracing_rest_dev?sslmode=disable"
ENV TRACKING_REST_PORT="3000"
ENTRYPOINT ["/tracing-rest"]
