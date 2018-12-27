package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "pgpool2"
)

var (
	NodeCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_count"),
		"How many nodes are in the pool at the moment",
		nil, nil,
	)
)

func InitMetricsExporter(pgpoolClient *PgPool) {
	exporter := MetricsExporter{pgpoolClient}
	prometheus.MustRegister(exporter)
}

type MetricsExporter struct {
	pgpoolClient *PgPool
}

func (exporter MetricsExporter) Collect(channel chan<- prometheus.Metric) {
	client := exporter.pgpoolClient
	errors := []error{}

	if nodeCount, err := client.GetNodeCount(); err == nil {
		channel <- prometheus.MustNewConstMetric(
			NodeCount, prometheus.GaugeValue, float64(nodeCount),
		)
	} else {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		log.Printf("Got errors when collecting metrics: %s\n", errors)
	}
}

func (exporter MetricsExporter) Describe(channel chan<- *prometheus.Desc) {
	channel <- NodeCount
}
