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
