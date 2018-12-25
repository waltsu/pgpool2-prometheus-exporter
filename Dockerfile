FROM golang:1.10-stretch as builder

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ADD . /go/src/github.com/waltsu/pgpool2-prometheus-exporter/

WORKDIR /go/src/github.com/waltsu/pgpool2-prometheus-exporter

RUN dep ensure
RUN go build .

FROM debian:jessie-slim

# Install pgpool to get pcp_* binaries
RUN \
  apt-get update && \
  apt-get --assume-yes --no-install-recommends install \
    curl \
    ca-certificates && \
  curl -L https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
  sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ jessie-pgdg main" >> /etc/apt/sources.list.d/pgdg.list' && \
  apt-get update && \
  apt-get --assume-yes --no-install-recommends install \
    pgpool2=4.0.2-1.pgdg80+1 && \
  rm -rf /var/lib/apt/lists/*


RUN mkdir /app
COPY --from=builder /go/src/github.com/waltsu/pgpool2-prometheus-exporter/pgpool2-prometheus-exporter /app/
COPY --from=builder /go/src/github.com/waltsu/pgpool2-prometheus-exporter/entrypoint.sh /app/
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]
CMD ["./pgpool2-prometheus-exporter"]
