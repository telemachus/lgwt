
+ You can create functions inside functions, but only if you assign an anonymous function to a variable. E.g., `assertCorrectMessage := func(t testing.TB, got, want string)`.
+ When you create a helper function for tests, you should accept a [`testing.TB`][TB] instead of `*testing.T` or `*testing.B`. `testing.TB` is an interface that both `*testingT` and `*testing.B` satisfy. If you accept the interface, rather than a specific pointer, you can use the helper functions in both tests and benchmarks.
+ Helper functions should call [`t.Helper` or `b.Helper`][Helper]. This allows the test runner to properly report line numbers. You can safely call `Helper` from multiple goroutines.

The testing cycle, according to James:

1. Write a test.
1. Make the compiler pass.
1. Run the test, confirm that it fails and that it has a meaningful error message.
1. Write enough code to make the test pass.
1. Refactor.

He argues that this discipline “may seem tedious but sticking to the feedback loop is important.”

[TB]: https://golang.org/pkg/testing/#TB
[Helper]: https://golang.org/pkg/testing/#T.Helper
