package main

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

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
