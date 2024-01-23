package main

import (
	"log"
	"net/http"

	"github.com/arturbaccarin/learn-go-with-tests/build-an-application/http-server/server"
)

func main() {
	server := &server.PlayerServer{Store: server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
