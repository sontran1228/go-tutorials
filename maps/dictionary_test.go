package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)

	})

	t.Run("unknow word", func(t *testing.T) {

		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)

	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s given, %s", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
