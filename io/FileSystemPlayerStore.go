package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore is to store data in file
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague return all players information
func (s *FileSystemPlayerStore) GetLeague() League {
	s.database.Seek(0, 0)
	league, _ := NewLeague(s.database)
	return league
}

// GetPlayerScore receive player name, then return player score
func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := s.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin record the win player
func (s *FileSystemPlayerStore) RecordWin(name string) {
	league := s.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	s.database.Seek(0, 0)
	json.NewEncoder(s.database).Encode(league)
}
