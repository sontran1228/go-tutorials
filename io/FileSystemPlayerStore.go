package main

import (
	"io"
)

// FileSystemPlayerStore is to store data in file
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague return all players information
func (s *FileSystemPlayerStore) GetLeague() []Player {
	s.database.Seek(0, 0)
	league, _ := NewLeague(s.database)
	return league
}

// GetPlayerScore receive player name, then return player score
func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range s.GetLeague() {
		if name == player.Name {
			wins = player.Wins
			break
		}
	}
	return wins

}
