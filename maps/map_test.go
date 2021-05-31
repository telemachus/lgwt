package dictionary

import "testing"

func TestSearchNaive(t *testing.T) {
	m := map[string]string{"test": "this is just a test"}

	actual := Search(m, "test")
	expected := "this is just a test"

	assertStrings(t, actual, expected, m)
}

func TestSearchWithCustomType(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		actual, _ := d.Search("test")
		expected := "this is just a test"

		assertStrings(t, actual, expected, d)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		k := "test"
		v := "this is just a test"
		err := d.Add(k, v)

		assertError(t, err, nil)
		assertDefinition(t, d, k, v)
	})

	t.Run("existing word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"
		d := Dictionary{k: v}
		err := d.Add(k, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, k, v)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"
		d := Dictionary{k: v}
		newV := "new definition"

		err := d.Update(k, newV)

		assertError(t, err, nil)
		assertDefinition(t, d, k, newV)
	})

	t.Run("new word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"
		d := Dictionary{}

		err := d.Update(k, v)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	k := "test"
	d := Dictionary{k: "test definition"}

	d.Delete(k)

	_, err := d.Search(k)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", k)
	}
}

func assertStrings(t testing.TB, actual string, expected string, m map[string]string) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %q; actual %q; given %#v", expected, actual, m)
	}
}

func assertError(t testing.TB, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual %q; expected %q", actual, expected)
	}
}

func assertDefinition(t testing.TB, d Dictionary, k, expected string) {
	t.Helper()

	actual, err := d.Search(k)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if actual != expected {
		t.Errorf("actual %q; expected %q", actual, expected)
	}
}
