package maps

// Dictionary is a map to store key:value
// - Maps is reference type therefore we can modify them without passing them as a pointer
// - Becuase maps is reference type, so maps can be a nil value:
//     + A nil map behaves like an empty map when reading,
//     + BUT attempts to write to a nil map will cause a runtime panic
// - Therefore, you should never initialize an empty map variable
//     + var m map[string]string               (X)
//     + dictionary = map[string]string{}      (V)
//     + dictionary = make(map[string]string)  (V)
type Dictionary map[string]string

// DictionaryErr is used to hold error message
type DictionaryErr string

const (
	// ErrNotFound is definition of an error in case the word couldn't be found in the dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists is definition of an error in case the word have been already existing in the dictionary
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	// ErrWordDoesNotExist is definition of an error in case the word haven't been existing in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// implemented error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

// Search is used to find a value from a "word"
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add is used to put a pair of word:definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update is used to update the existing definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete is used to delete a definition
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
