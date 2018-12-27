package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	nodeCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "node_count",
		Help: "How many nodes are in the pool at the moment",
	})
)

func ExportMetrics(metricsChannel <-chan Metrics) {
	for {
		metrics := <-metricsChannel
		log.Println("Metrics exported")

		nodeCountGauge.Set(float64(metrics.nodeCount))
	}

}

func init() {
	prometheus.MustRegister(nodeCountGauge)
}
