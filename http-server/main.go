package main

import (
	"log"
	"net/http"
)

// NewInMemoryPlayerStore init NewInMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore is an implementation of PlayerStore to hold player information in Memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore receive player name, then return player score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin record the win player
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
