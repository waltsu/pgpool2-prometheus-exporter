package main

import (
  "log"
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
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
  log.Println("Gathering metrics")

  bashExecutor := new(BashExecutor)
  pgpool := NewPgPool(bashExecutor)

  gatherMetrics(pgpool)
}

func gatherMetrics(pgpool *PgPool) {
  nodeCount, err := pgpool.GetNodeCount()

  if err != nil {
    log.Fatal(err);
    return
  }
  log.Println(nodeCount)
}
