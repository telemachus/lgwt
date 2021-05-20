# Hello, World

James begins with several simple “Hello, world!” functions. He uses these examples to demonstrate the basics of Go and Go’s testing package.

## Testing in Go

How do you write tests in Go? Simple: you import testing from Go’s standard library, and then you write code like this:

```go
func TestSomeNameHere(t *testing.T) {
    got := SomeFunctionHere()
    want := "some return value here"

    if got != want {
        t.Errorf("got %q; want %q", got, want)
    }
}
```

+ Tests belong in files with names like `xxx_test.go`. Go’s tools will treat such files in special ways.
+ Test functions must start with the word `Test`.
+ Test functions take only one argument, a pointer to `testing.T`.
+ If you want to use `testing.T`, you must import testing.
+ The pointer `t` gives your testing code access to methods like `Errorf`. To see all the methods on `t`, see [the fine manual[testing-t].

[testing-t]: https://golang.org/pkg/testing/#T
