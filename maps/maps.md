# *Learn Go With Tests*, Maps

Maps are largely what you would expect. You store items with a key, and as a result, you can find an item in a map much more easily than in an array or slice. You can use any [comparable type][comparable] as the key.

You can create a map in several ways.

+ You can use a map literal: `d := map[string]string{"foo": "bar"}`
+ With an empty map literal: `d := map[string]string{}`
+ With `make`: `d := make(map[string]string)`.

Beware, however, not to initialize an empty map variable. That is, don’t do this: `m := map[string]string`. That map is `nil`, and if you try to add something to an `nil` map, Go will throw a panic at runtime. (Why the hell does the language allow this? When is it useful?)

Although you can’t tell by looking at their declaration or use, Go’s maps are pointers to an underlying datatype. Thus, you *can* change a map that you pass to a function. (This is a sharp contrast to structs where you must carefully consider whether you need to modify them and then manually use pointers accordingly.)

[comparable]: https://golang.org/ref/spec#Comparison_operators

## Adding and Deleting

James recommends that you create two methods `Add` and `Update` to distinguish between, well, adding a new entry and updating an existing entry. If you use `Add` on an entry that already exists, you get an error. If you use `Update` on an entry that does not yet exist, you get an error. This is probably best since you don’t want to accidentally update when you mean to add a new entry.

On the other hand, there seems to be no need to signal an error if you try to delete an entry that does not exist. Why not? Well, if the end result you want is no such entry, then you reach your goal either way. No harm is done if there was never such an entry to begin with. (Are there exceptions to this? Probably, but I am not sure I need to worry about that now.)

## Moar Errors!

James creates a more complicated error structure late in this chapter. In an earlier chapter, he began to use global variables as errors, but here he switches to actual constants.

```go
const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}
```

James points us to [an article by Dave Cheney][constant-errors] for more discussion and explanation. 

[constant-errors]: https://dave.cheney.net/2016/04/07/constant-errors
