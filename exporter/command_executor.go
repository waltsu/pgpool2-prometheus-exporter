package main

import (
  "bytes"
)

type CommandExecutor interface {
  Execute(command string, args ...string) (*bytes.Buffer, error)
}

type BashExecutor struct {}

func (bash *BashExecutor) Execute(command string, args ...string) (*bytes.Buffer, error) {
  return bytes.NewBufferString("test"), nil
}
