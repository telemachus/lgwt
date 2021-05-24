package main

import (
	"fmt"
	"testing"
)

func assertCorrectGreeting(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected: %q; actual: %q", expected, actual)
	}
}

func TestHello(t *testing.T) {
	testWithEmptyString := func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World!"
		assertCorrectGreeting(t, actual, expected)
	}

	testWithName := func(t *testing.T) {
		actual := Hello("Peter", "")
		expected := "Hello, Peter!"
		assertCorrectGreeting(t, actual, expected)
	}

	testInSpanish := func(t *testing.T) {
		actual := Hello("Bluie", "sp")
		expected := "¡Hola, Bluie!"
		assertCorrectGreeting(t, actual, expected)
	}

	testInFrench := func(t *testing.T) {
		actual := Hello("Chef", "fr")
		expected := "Bonjour, Chef!"
		assertCorrectGreeting(t, actual, expected)
	}

	t.Run("empty string", testWithEmptyString)
	t.Run("with name", testWithName)
	t.Run("in Spanish", testInSpanish)
	t.Run("in French", testInFrench)
}

// You can put testable examples into your code. I think that this is meant
// for documentation rather than test files, but maybe I’m wrong?
func ExampleHello() {
	fmt.Println(Hello("Peter", "en"))
	// Output: Hello, Peter!

	fmt.Println(Hello("", "en"))
	// Output: Hello, World!
}
