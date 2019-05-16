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
	http.Handler
}

// NewPlayerServer init PlayerServer object. It's one time initialization of router
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router
	return p
}

// PlayerServer return player information.
// This function is no longer needed, because the http.Hanlder is already embedded to PlayerServer.
// (refer to https://golang.org/doc/effective_go.html#embedding)
// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.ServeHTTP(w, r)
// }

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
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
