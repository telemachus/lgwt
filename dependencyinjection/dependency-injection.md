# *Learn Go With Tests*, Dependency Injection

What is dependency injection? Dependency injection is when code passes a dependency into a function or method rather than assuming a fixed dependency internally. The passing is the injection. Here’s a brief example of code without dependency injection.

```go
func Greet(name string) {
    fmt.Printf("Hello, %s!", name)
}
```

The function `Greet` is difficult to test because output goes immediately to standard output via `Printf`. We can make the function more flexible—and more testable—if we explicitly tell the function where to send its output.

```go
func Greet(w io.Writer, name string) {
    fmt.Fprintf(w, "Hello, %s!", name)
}

// Code that tests Greet looks like this.
func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

// Code that actually uses Greet looks like this.
Greet(os.Stdout, "World")
```
