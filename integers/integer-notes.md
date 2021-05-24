# *Learn Go With Tests*: Integers

## Testable Examples

You can include testable examples as part of your tests. I mentioned this earlier, but at the time, I thought they were mostly for documentation. I was confused.

Testable examples are for tests. You need to print the function’s return values to standard out in order to test with `Output`. It’s a bit clunky, but it works. Because Go automatically runs examples in test files, you know that the examples and code cannot get out of sync. That is, if you run your tests, and your examples are no longer valid, they will fail like tests will fail if the underlying code no longer works correctly.
