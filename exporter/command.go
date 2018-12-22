package main

import (
  "bytes"
)

type Command interface {
  Execute(command string, args ...string) (*bytes.Buffer, error)
}

type BashCommand struct {}

func (bash *BashCommand) Execute(command string, args ...string) (*bytes.Buffer, error) {
  return bytes.NewBufferString("test"), nil
}
