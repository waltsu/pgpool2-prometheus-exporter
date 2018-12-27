package exporter

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	PcpLocation = GetEnv("PCP_LOCATION", "/usr/sbin/")

	PcpUsername = GetEnv("PCP_USER", "pcpuser")
	PcpHost     = GetEnv("PCP_HOST", "localhost")
	PcpPort     = GetEnv("PCP_PORT", "9898")

	PcpDefaultArguments = []string{"--username=" + PcpUsername, "--host=" + PcpHost, "--port=" + PcpPort, "-w"}
)

type PgPool struct {
	commandExecutor CommandExecutor
}

func NewPgPool(executor CommandExecutor) *PgPool {
	log.Println(fmt.Sprintf("Create new pgpool client with location %s and default arguments %s", PcpLocation, PcpDefaultArguments))
	pgpool := &PgPool{executor}
	return pgpool
}

func (pgpool PgPool) StartMetricsProducer(metricsChannel chan<- *Metrics) {
	for {
		metrics := &Metrics{}

		if nodeCount, err := pgpool.GetNodeCount(); err == nil {
			metrics.nodeCount = nodeCount
		}

		metricsChannel <- metrics
		log.Println("Metrics gathered")

		time.Sleep(1 * time.Second) // TODO: Set sleep value from env variable
	}
}

func (pgpool PgPool) GetNodeCount() (int, error) {
	response, err := pgpool.commandExecutor.Execute(PcpLocation+"pcp_node_count", PcpDefaultArguments...)
	if err != nil {
		return -1, err
	}

	node_count, err := strconv.Atoi(strings.TrimSuffix(response.String(), "\n"))
	if err != nil {
		return -1, err
	}

	return node_count, nil
}
