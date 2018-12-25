package exporter

import (
	"bytes"
	"os/exec"
)

type CommandExecutor interface {
	Execute(command string, args ...string) (*bytes.Buffer, error)
}

type BashExecutor struct{}

func (bash *BashExecutor) Execute(command string, args ...string) (*bytes.Buffer, error) {
	execCommand := exec.Command(command, args...)

	stdout := &bytes.Buffer{}
	execCommand.Stdout = stdout
	execCommand.Stderr = stdout

	error := execCommand.Run()
	if error != nil {
		return nil, error
	}
	return stdout, nil
}

type TestExecutor struct{}

func (bash *TestExecutor) Execute(command string, args ...string) (*bytes.Buffer, error) {
	return bytes.NewBufferString("test"), nil
}
