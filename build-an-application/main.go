package main

import (
	"log"
	"net/http"

	"github.com/arturbaccarin/learn-go-with-tests/build-an-application/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
