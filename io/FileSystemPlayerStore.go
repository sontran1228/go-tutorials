package main

import (
	"encoding/json"
	"os"
)

// FileSystemPlayerStore is to store data in file
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore init FileSystemPlayerStore
func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

// GetLeague return all players information
func (s *FileSystemPlayerStore) GetLeague() League {
	return s.league
}

// GetPlayerScore receive player name, then return player score
func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := s.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin record the win player
func (s *FileSystemPlayerStore) RecordWin(name string) {
	player := s.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{name, 1})
	}

	s.database.Encode(s.league)
}
