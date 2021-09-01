package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http:/localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentaion(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentaion",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func main() {
	http.HandleFunc("/", documentaion)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
