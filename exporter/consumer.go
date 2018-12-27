package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "pgpool2"
)

var (
	/*
	nodeCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "node_count",
		Help: "How many nodes are in the pool at the moment",
	})
	*/
	NodeCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_count"),
		"How many nodes are in the pool at the moment",
		nil, nil,
	)
)

func NewConsumer() MetricsConsumer {
	consumer := MetricsConsumer{}

	prometheus.MustRegister(consumer)

	return consumer
}

func (consumer *MetricsConsumer) StartMetricsConsumer(metricsChannel <-chan *Metrics) {
	log.Println("Starting metrics consumer")
	for {
		consumer.lastGatheredMetrics = <- metricsChannel
	}

}

type MetricsConsumer struct {
	lastGatheredMetrics *Metrics
}

func (consumer MetricsConsumer) Collect(channel chan<- prometheus.Metric) {
}

func (consumer MetricsConsumer) Describe(channel chan<- *prometheus.Desc) {
	channel <- NodeCount
}
