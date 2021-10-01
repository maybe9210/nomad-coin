package main

import (
	"github.com/maybe9210/nomad-coin/blockchain"
	"github.com/maybe9210/nomad-coin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
