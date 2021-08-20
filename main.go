package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/maybe9210/nomad-coin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockChain().GetAllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
