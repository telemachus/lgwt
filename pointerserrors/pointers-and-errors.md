# *Learn Go With Tests*, Pointers and Errors

## Pointers

When you use structs, you will probably want to store state in them. Imagine you have a structure that represents a bank account. You need to allow the user to deposit and withdraw money from the account. In order to do that, the user must be able to change the state of the account balance, which the struct stores.

No problem, right? Well, sort of. You can absolutely do what you need, but you must use pointers. If you pass a struct itself into a function, the struct is copied. When you change the copy, those changes are not reflected in the original. Thus, the naive method won’t work. Instead, you need to define a method on a pointer to struct.

```go
func (w Wallet) Deposit(amt float64) {
    w.balance += amt // Nope: this won’t work.
}

func (w *Wallet) Deposit(amt float64) {
    w.balance += amt // Yup: this works.
}
```

Notice that in the second case we don’t explicitly dereference the value of `w`. You can explicitly dereference the pointer, and it would look like this: `(*w).balance += amt`. That code works. However, [Go assumes that you mean to dereference a struct pointer][mvalues]. In this case, the code does what you mean not (precisely) what you say.

By convention methods on a receiver should all take values or all take pointers.

Even if you are not going to mutate an argument, you may want to pass a pointer rather than copy a value. Look out for especially large values or singletons.

### `nil` and Pointers

+ Pointers can be nil.
+ If a function returns a pointer, you should make sure that it’s not `nil`. If you try to call a method on a `nil` pointer, you will get a panic *at runtime*. The compiler won’t catch this in advance since the compiler cannot know when a pointer is nil before the program runs.
+ Returning `nil` can be useful to signal a missing value.

[mvalues]: https://golang.org/ref/spec#Method_values

## Custom Types

You can create custom types based on existing type to make your code clearer. For example, if your wallet stores bitcoin, you can declare a custom type for it.

```go
type Bitcoin int
```

Bitcoin will support all the built-in operations of the underlying type (here integers), but your code will be clearer.

Also, once you have a custom method, you can declare methods on that type.

```go
func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
```

Now your custom type has a default string representation. This makes your type much easier to use with `fmt` functions.

## Errors

Your functions will often need to signal that there is a problem or the user has done something wrong. For example, if you have a method to withdraw money from a bank account, a user may try to withdraw more money than is in the account.

You use `errors.New(<some string goes here>)` to indicate an error. It’s up to the calling code to check for the error and handle it accordingly. (Note that I should use `errcheck` to verify that I’ve checked and handled all the errors I should be checking and handling.)

### Testing and Errors

You should define a global error to test for rather than testing for specific strings. (This helps you to avoid the case where the error changes in the module but not in the tests.)
