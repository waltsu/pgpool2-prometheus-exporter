package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/waltsu/pgpool2-prometheus-exporter/exporter"
)

func main() {
	go startMetricGathering()
	startPrometheusServer()
}

func startPrometheusServer() {
	log.Println("Starting Prometheus HTTP server")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startMetricGathering() {
	log.Println("Start gathering metrics")

	commandExecutor := new(exporter.BashExecutor)
	pgpool := exporter.NewPgPool(commandExecutor)
	consumer := exporter.NewConsumer()
	metricsChannel := make(chan *exporter.Metrics)

	go pgpool.StartMetricsProducer(metricsChannel)
	go consumer.StartMetricsConsumer(metricsChannel)
}
