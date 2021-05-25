# *Learn Go With Tests*, Arrays and Slices

## Arrays

Arrays are rarely used in Go. Why? Because their size cannot be changed, it must be declared in advance, and it is part of the type of the array. If you write a function to sum items in an array, it will only work on arrays of one size.

You can declare arrays in two ways:

+ `[N]type{value0, value1, …, valueN-1}`
+ `[...]type{value0, value1, …, valueN-1}`

In the second case, the compiler counts the number of items in the array for you. That might sound convenient, but the convenience is minimal. You must still know what the number is if you want to use that array in functions.
