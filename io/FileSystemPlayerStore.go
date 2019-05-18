package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// FileSystemPlayerStore is to store data in file
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore init FileSystemPlayerStore
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
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
