package dictionary

import "testing"

func TestFind(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dict.Find("test")
		want := "this is just a test"

		compareStrings(t, got, want)
	})

	t.Run("unknow world", func(t *testing.T) {
		_, got := dict.Find("unknow")

		compareErr(t, got, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dict.Add(word, definition)

		compareErr(t, err, nil)
		compareDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "new test")

		compareErr(t, err, ErrExistingWord)
		compareDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dict := Dictionary{word: definition}
		err := dict.Update(word, newDefinition)

		compareErr(t, err, nil)
		compareDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, definition)

		compareErr(t, err, ErrNonExistingWord)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "test definition"}

		dict.Delete(word)

		_, err := dict.Find(word)
		if err != ErrWordNotFound {
			t.Errorf("'%s' is expected to be deleted", word)
		}
	})
}

func compareStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s', given '%s'", got, want, "test")
	}
}

func compareErr(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func compareDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Find(word)
	if err != nil {
		t.Fatal("should have returned a word: ", err)
	}

	if definition != got {
		t.Errorf("got '%s', want '%s'", got, definition)
	}
}
