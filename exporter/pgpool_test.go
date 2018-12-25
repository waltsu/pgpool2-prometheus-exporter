package exporter

import (
	"bytes"
	"errors"

  "testing"
  "github.com/stretchr/testify/assert"
)

type TestExecutor struct{
	returnStdout *bytes.Buffer
	errorValue error
}

func NewTestExecutor(stdout string, errorValue error) TestExecutor{
	return TestExecutor{bytes.NewBufferString(stdout), errorValue}
}

func (executor TestExecutor) Execute(command string, args ...string) (*bytes.Buffer, error) {
	return executor.returnStdout, executor.errorValue
}

func TestGetNodeCountReturnsNodeCount(t *testing.T) {
	testExecutor := NewTestExecutor("5\n", nil)
	pgpool := NewPgPool(testExecutor)

	nodeCount, _ := pgpool.GetNodeCount()
	assert.Equal(t, nodeCount, 5)
}

func TestGetNodeCountFailsWhenInvocationFails(t *testing.T) {
	testExecutor := NewTestExecutor("", errors.New("boom"))

	pgpool := NewPgPool(testExecutor)

	nodeCount, err := pgpool.GetNodeCount()
	assert.Equal(t, nodeCount, -1)
	assert.NotNil(t, err);
}

func TestGetNodeCountFailsWithMalformedStdout(t *testing.T) {
	testExecutor := NewTestExecutor("foobar", nil)

	pgpool := NewPgPool(testExecutor)

	nodeCount, err := pgpool.GetNodeCount()
	assert.Equal(t, nodeCount, -1)
	assert.NotNil(t, err);
}
