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

## Various

+ You can create functions inside functions, but only if you assign an anonymous function to a variable. E.g., `assertCorrectMessage := func(t testing.TB, got, want string)`.
+ When you create a helper function for tests, you should accept a [`testing.TB`][TB] instead of `*testing.T` or `*testing.B`. `testing.TB` is an interface that both `*testingT` and `*testing.B` satisfy. If you accept the interface instead of a specific pointer, you can use the helper functions in both tests and benchmarks.
+ Helper functions should call [`t.Helper` or `b.Helper`][Helper]. This allows the test runner to properly report line numbers. You can safely call `Helper` from multiple goroutines.

## The Discipline of Testing

The testing cycle, according to James:

1. Write a test for a function that does not even exist. The program should not compile.
1. Make the compiler pass by writing the outline of the function, but not a proper body.
1. Run the test, confirm that it fails and that it has a meaningful error message.
1. Write enough code to make the test pass.
1. Refactor.

He argues that this discipline “may seem tedious but sticking to the feedback loop is important.”

[TB]: https://golang.org/pkg/testing/#TB
[Helper]: https://golang.org/pkg/testing/#T.Helper
