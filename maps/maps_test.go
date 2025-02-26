package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known key", func(t *testing.T) {
		got, _ := Dictionary.Search(dictionary, "test")
		want := "this is a test"
		assertString(t, got, want, dictionary)

	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := Dictionary.Search(dictionary, "unknown")
		want := ErrNotFound
		if err == nil {
			t.Fatal("Expected to get an error.")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("add a word", func(t *testing.T) {
		dictionary.Add("test", "this is a test")
		want := "this is a test"
		got, err := dictionary.Search("test")
		if err == nil {
			assertString(t, got, want, dictionary)
		}

	})

	t.Run("add an existing word", func(t *testing.T) {
		key := "test existing"
		value := "this is a test value"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)

	})
}

func assertString(t testing.TB, got, want string, dictionary map[string]string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q and wanted %q, dictionary is %q", got, want, dictionary)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q but want error %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatalf("wanted word %q but found error %q", value, err)
	}

	assertString(t, got, value, dictionary)
}
