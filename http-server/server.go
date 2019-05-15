package main

import (
	"fmt"
	"net/http"
)

// PlayerStore is an interface that acts as repository to store player score
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information
// It will an implementation of "http.Handler" interface by implementing function "ServeHTTP"
// see more: server.go -> type Handler interface
type PlayerServer struct {
	store PlayerStore
}

// PlayerServer return player information.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {

	p.store.RecordWin(player)

	w.WriteHeader(http.StatusAccepted)
}
