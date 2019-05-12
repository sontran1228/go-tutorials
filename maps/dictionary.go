package maps

import "errors"

// Dictionary is a map to store key:value
type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

// Search is used to find a value from a "word"
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
