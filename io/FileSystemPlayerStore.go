package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore is to store data in file
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

// NewFileSystemPlayerStore init FileSystemPlayerStore
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
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

	s.database.Seek(0, 0)
	json.NewEncoder(s.database).Encode(s.league)
}
