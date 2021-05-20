package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := Hello("Peter")
	want := "Hello, Peter!"

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

// You can put testable examples into your code. I think that this is meant
// for documentation rather than test files, but maybe I’m wrong?
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
