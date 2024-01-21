package server

import (
	"fmt"
	"net/http"
	"strings"
)

// It doesn't feel right that our server knows the scores.
/*
We moved the score calculation out of the main body of our handler into a function GetPlayerScore.
This feels like the right place to separate the concerns using interfaces.
*/
// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
