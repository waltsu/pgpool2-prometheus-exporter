package exporter

import (
	"bytes"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

type TestExecutor struct {
	returnStdouts []*bytes.Buffer
	errorValue   error
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
	if len(executor.returnStdouts) < 2 {
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
