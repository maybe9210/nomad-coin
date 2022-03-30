package main

import (
	"github.com/maybe9210/nomad-coin/cli"
	"github.com/maybe9210/nomad-coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
	// wallet.Start()
	// wallet.Wallet()
}
