package exporter

import (
	"log"
)

var (
	PcpLocation = GetEnv("PCP_LOCATION", "/usr/sbin/")

	PcpUsername = GetEnv("PCP_USERNAME", "pcpuser")
	PcpHost = GetEnv("PCP_HOST", "localhost")
	PcpPort = GetEnv("PCP_PORT", "9898")

	PcpDefaultArguments = []string{ "--username=" + PcpUsername, "--host=" + PcpHost, "--port=" + PcpPort, "-w" }
)

type PgPool struct {
	commandExecutor CommandExecutor
}

func NewPgPool(executor CommandExecutor) *PgPool {
	pgpool := &PgPool{executor}
	return pgpool
}

func (pgpool *PgPool) GetNodeCount() (int, error) {
	response, err := pgpool.commandExecutor.Execute(PcpLocation + "pcp_node_count", PcpDefaultArguments...)
	if err != nil {
		return -1, err
	}
	log.Println(response.String())

	return 0, nil
}
