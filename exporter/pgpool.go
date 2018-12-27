package exporter

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
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

type NodeInfo struct {
	hostname         string
	port             int
	status           int
	weight           float64
	role             string
	replicationDelay int
}

func NewPgPool(executor CommandExecutor) *PgPool {
	log.Println(fmt.Sprintf("Create new pgpool client with location %s and default arguments %s", PcpLocation, PcpDefaultArguments))
	pgpool := &PgPool{executor}
	return pgpool
}

func (pgpool PgPool) GetNodeCount() (int, error) {
	response, err := pgpool.commandExecutor.Execute(PcpLocation+"pcp_node_count", PcpDefaultArguments...)
	if err != nil {
		return -1, err
	}

	node_count, err := strconv.Atoi(trimNewLine(response))
	if err != nil {
		return -1, err
	}

	return node_count, nil
}

func (pgpool PgPool) GetNodeInfos() []NodeInfo {
	nodeInfos := []NodeInfo{}

	nodeCount, err := pgpool.GetNodeCount()
	if err != nil {
		log.Println(err)
		return nodeInfos
	}

	nodeInfoRegexp := regexp.MustCompile(`^(.*?)\s(\d*?)\s(.\d*?)\s(.*?)\s(.*?)\s(.*?)\s(\d*?)\s.*$`)
	for i := 0; i < nodeCount; i++ {
		argumentsWithNodeIndex := append(PcpDefaultArguments, string(i))
		response, err := pgpool.commandExecutor.Execute(PcpLocation+"pcp_node_info", argumentsWithNodeIndex...)
		if err != nil {
			log.Println(err)
			return nodeInfos
		}

		trimmedResponse := trimNewLine(response)
		rawNodeInfo := nodeInfoRegexp.FindStringSubmatch(trimmedResponse)
		nodeInfos = append(nodeInfos, buildNodeInfo(rawNodeInfo))
	}
	return nodeInfos
}

// Strips possible new line from the end of line and returns it as a string
func trimNewLine(line *bytes.Buffer) string {
	return strings.TrimSuffix(line.String(), "\n")
}

func buildNodeInfo(rawNodeInfo []string) NodeInfo {
	nodeInfo := NodeInfo{}

	if len(rawNodeInfo) != 8 {
		log.Printf("Wrong amount of string elements: '%s'", rawNodeInfo)
		return nodeInfo
	}

	nodeInfo.hostname = rawNodeInfo[1]
	nodeInfo.port, _ = strconv.Atoi(rawNodeInfo[2])
	nodeInfo.status, _ = strconv.Atoi(rawNodeInfo[3])
	nodeInfo.weight, _ = strconv.ParseFloat(rawNodeInfo[4], 64)
	nodeInfo.role = rawNodeInfo[6]
	if replicationDelay, err := strconv.Atoi(rawNodeInfo[7]); err != nil {
		nodeInfo.replicationDelay = -1
	} else {
		nodeInfo.replicationDelay = replicationDelay
	}

	return nodeInfo
}
