package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"testing": "just a test"}

	t.Run("known", func(t *testing.T) {
		got, _ := dictionary.Search("testing")
		//got := dictionary["testing"]
		want := "just a test"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("unkown", func(t *testing.T) {
		_, err := dictionary.Search("haha")
		if err != NotFoundErr {
			t.Errorf("got '%s' want '%s'", err, NotFoundErr)
		}
	})

	t.Run("add", func(t *testing.T) {
		dictionary.Add("add", "added")
		got, _ := dictionary.Search("add")
		want := "added"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
