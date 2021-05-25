# *Learn Go With Tests*, Arrays and Slices

## Arrays

Arrays are rarely used in Go. Why? Because their size cannot be changed, it must be declared in advance, and it is part of the type of the array. If you write a function to sum items in an array, it will only work on arrays of one size.

You can declare arrays in two ways:

+ `[N]type{value0, value1, …, valueN-1}`
+ `[...]type{value0, value1, …, valueN-1}`

In the second case, the compiler counts the number of items in the array for you. That might sound convenient, but the convenience is minimal. You must still know what the number is if you want to use that array in functions.

If you want a function to accept slices and arrays, you can use a variadic argument and manually unpack the input.

```go
func Sum(ns ...int) int {
	var total int
	for _, n := range ns {
		total += n
	}

	return total
}

// Call it one of the following ways.
ns := [5]int{1,2,3,4,5}
sn := []int{1,2,3,4,5}
foo := Sum(ns[:]...)
bar := Sum(sn...)
bizz := Sum(1,2,3,4,5)
```

The `...` goes *before* the type in a function declaration, but *after* the argument in a function call. When `...` appears before something, it is the *pack operator*. When `...` appears after something, it is the *unpack operator*.

You cannot use `...` directly on an array. First, you have to convert the array into a slice using `[:]`.
