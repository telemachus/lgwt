package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("with name", func(t *testing.T) {
		b := bytes.Buffer{}
		Greet(&b, "Chris")
		
		actual := b.String()
		expected := "Hello, Chris!"
		
		if actual != expected {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})

	t.Run("without name", func(t *testing.T) {
		b := bytes.Buffer{}
		Greet(&b, "")

		actual := b.String()
		expected := "Hello, World!"

		if actual != expected {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})
}
