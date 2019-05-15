package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore is an implementation of PlayerStore to hold player information in Memory
type InMemoryPlayerStore struct{}

// GetPlayerScore receive player name, then return player score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin record the win player
func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
