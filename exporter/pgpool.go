package exporter

import (
	"log"
)

const (
	PcpLocation = "/usr/sbin/" // TODO: Configure from env variables
)

type PgPool struct {
	commandExecutor CommandExecutor
}

func NewPgPool(executor CommandExecutor) *PgPool {
	pgpool := &PgPool{executor}
	return pgpool
}

func (pgpool *PgPool) GetNodeCount() (int, error) {
	response, err := pgpool.commandExecutor.Execute(PcpLocation + "pcp_node_count")
	if err != nil {
		return -1, err
	}
	log.Println(response.String())

	return 0, nil
}
