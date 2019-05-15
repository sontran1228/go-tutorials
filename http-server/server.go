package main

import (
	"fmt"
	"net/http"
)

// PlayerStore is an interface that acts as repository to store player score
type PlayerStore interface {
	GetPlayerScore(name string) int
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
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
