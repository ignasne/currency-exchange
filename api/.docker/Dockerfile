FROM golang:1.12-alpine as builder

ENV CGO_ENABLED=0

RUN apk add --update git && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /opt/source
WORKDIR /opt/source

ADD go.mod go.sum ./
RUN go mod vendor

ADD ./ ./

RUN go build -o ./bin/quotes-api .

FROM alpine:3.9

ENV APP_MIGRATIONS_PATH /etc/quotes-api/migrations
ENV RATES_API_URL https://api.exchangeratesapi.io

RUN apk --no-cache add tzdata

COPY --from=builder /opt/source/bin/quotes-api /usr/local/bin/quotes-api
COPY .docker/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
COPY migrations/* /etc/quotes-api/migrations/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["quotes-api", "server"]
