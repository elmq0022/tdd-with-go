package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		key := "test"

		got, err := dictionary.Search(key)
		_ = err

		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		key := "unknown"

		_, err := dictionary.Search(key)
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add a new word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("add an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {

		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("update nonexisting word", func(t *testing.T) {

		word := "test"
		newDefinition := "new definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q but want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, definition)
}
