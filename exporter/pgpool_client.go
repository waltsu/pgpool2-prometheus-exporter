package main

import (
  "log"
)

type Client struct {
  command Command
}

func PgPool(command Command) (*Client) {
  client := &Client{command}
  return client
}

func (client *Client) GetNodeCount() (int, error) {
  response, err := client.command.Execute("pcp_node_count")
  if err != nil {
    return -1, err;
  }
  log.Println(response.String())

  return 0, nil
}
