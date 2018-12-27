package exporter

import (
	"bytes"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

type TestExecutor struct {
	returnStdouts []*bytes.Buffer
	errorValue    error
}

func NewTestExecutor(stdouts []string, errorValue error) *TestExecutor {
	buffers := []*bytes.Buffer{}
	for _, stdout := range stdouts {
		buffers = append(buffers, bytes.NewBufferString(stdout))
	}
	return &TestExecutor{buffers, errorValue}
}

func (executor *TestExecutor) Execute(command string, args ...string) (*bytes.Buffer, error) {
	var returnStdout *bytes.Buffer
	if len(executor.returnStdouts) > 1 {
		returnStdout, executor.returnStdouts = executor.returnStdouts[0], executor.returnStdouts[1:]
	} else {
		returnStdout = executor.returnStdouts[0]
	}
	return returnStdout, executor.errorValue
}

func TestGetNodeCountReturnsNodeCount(t *testing.T) {
	testExecutor := NewTestExecutor([]string{"5\n"}, nil)
	pgpool := NewPgPool(testExecutor)

	nodeCount, _ := pgpool.GetNodeCount()
	assert.Equal(t, nodeCount, 5)
}

func TestGetNodeCountFailsWhenInvocationFails(t *testing.T) {
	testExecutor := NewTestExecutor([]string{""}, errors.New("boom"))

	pgpool := NewPgPool(testExecutor)

	nodeCount, err := pgpool.GetNodeCount()
	assert.Equal(t, nodeCount, -1)
	assert.NotNil(t, err)
}

func TestGetNodeCountFailsWithMalformedStdout(t *testing.T) {
	testExecutor := NewTestExecutor([]string{"foobar"}, nil)

	pgpool := NewPgPool(testExecutor)

	nodeCount, err := pgpool.GetNodeCount()
	assert.Equal(t, nodeCount, -1)
	assert.NotNil(t, err)
}

func TestGetNodeInfosReturnsInfoFromAllNodes(t *testing.T) {
	pcpStdouts := []string{
		"2\n",
		"postgres-1 5432 2 0.500000 up primary 0 2018-12-27 17:23:34\n",
		"postgres-2 5432 2 0.500000 up standby 23410272 2018-12-27 17:23:34\n",
	}
	testExecutor := NewTestExecutor(pcpStdouts, nil)

	pgpool := NewPgPool(testExecutor)

	nodeInfos := pgpool.GetNodeInfos()
	assert.Equal(t, 2, len(nodeInfos))

	firstInfo := nodeInfos[0]
	assert.Equal(t, "postgres-1", firstInfo.hostname)
	assert.Equal(t, 5432, firstInfo.port)
	assert.Equal(t, 2, firstInfo.status)
	assert.Equal(t, 0.5, firstInfo.weight)
	assert.Equal(t, "primary", firstInfo.role)
	assert.Equal(t, 0, firstInfo.replicationDelay)

	secondInfo := nodeInfos[1]
	assert.Equal(t, "postgres-2", secondInfo.hostname)
	assert.Equal(t, 5432, secondInfo.port)
	assert.Equal(t, 2, secondInfo.status)
	assert.Equal(t, 0.5, secondInfo.weight)
	assert.Equal(t, "standby", secondInfo.role)
	assert.Equal(t, 23410272, secondInfo.replicationDelay)
}

func TestGetNodeInfosReturnsEmptyNodeInfoWithMalformedStdout(t *testing.T) {
	pcpStdouts := []string{
		"1\n",
		"asdfasf\n",
	}
	testExecutor := NewTestExecutor(pcpStdouts, nil)

	pgpool := NewPgPool(testExecutor)

	nodeInfos := pgpool.GetNodeInfos()
	assert.Equal(t, "", nodeInfos[0].hostname)
	assert.Equal(t, 0, nodeInfos[0].port)
}
