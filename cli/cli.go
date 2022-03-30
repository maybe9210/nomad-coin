package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/maybe9210/nomad-coin/explorer"
	"github.com/maybe9210/nomad-coin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags! \n\n")
	fmt.Printf("-port:			Set the PORT of the server\n")
	fmt.Printf("-mode:			Choose one among 'html', 'rest', 'all'\n\n")

	// runtime.Goexit()
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set Port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port)
		second_port := *port + 1
		if second_port >= 10000 {
			second_port = *port - 1
		}
		explorer.Start(second_port)
	default:
		usage()
	}

}
