package main

import (
	"io"
)

// FileSystemPlayerStore is to store data in file
type FileSystemPlayerStore struct {
	database io.Reader
}

// GetLeague return all players information
func (s *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(s.database)
	return league
}
