FROM golang:alpine as builder

RUN apk add --no-cache curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ADD . /go/src/github.com/waltsu/pgpool2-prometheus-exporter/

WORKDIR /go/src/github.com/waltsu/pgpool2-prometheus-exporter

RUN dep ensure
RUN go build .

FROM alpine
RUN mkdir /app
COPY --from=builder /go/src/github.com/waltsu/pgpool2-prometheus-exporter/pgpool2-prometheus-exporter /app/
WORKDIR /app

EXPOSE 8080
CMD ["./pgpool2-prometheus-exporter"]
