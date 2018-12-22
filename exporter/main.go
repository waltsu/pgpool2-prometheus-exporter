package main

import (
  "log"
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Starting Prometheus HTTP server")
  http.Handle("/metrics", promhttp.Handler())
  log.Fatal(http.ListenAndServe(":8080", nil))

  gatherMetrics()
}

func gatherMetrics() {
  log.Println("Gathering metrics")
}
