package di

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) {
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}
