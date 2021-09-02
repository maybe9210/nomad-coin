package main

import (
	"github.com/maybe9210/nomad-coin/explorer"
	"github.com/maybe9210/nomad-coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
