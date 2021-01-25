package main

import (
	"log"
	"net/http"

	"github.com/aphiliplinell/go-hello"
)

const bindAddr = ":9191"

func main() {
	server := hello.Server{}
	router := hello.NewRouter(&server)

	log.Printf("Listening on %s", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, router))
}
