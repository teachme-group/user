FROM golang:1.23-alpine AS build
WORKDIR /app
ADD . /app

RUN apk --no-cache add gcc g++ make git \
    && go env -w  GOSUMDB=off \
    && go mod download \
    && go build -o /bin/main ./cmd/main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

COPY --from=build /bin/main /bin/main
COPY .env /app/.env

WORKDIR /app
ENTRYPOINT ["/bin/main"]