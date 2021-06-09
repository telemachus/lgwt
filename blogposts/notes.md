# Notes on “Reading Files”


## [Reading files](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#reading-files)

> We've been asked to create a package that converts a given folder of blog posts *and return* a collection of Post which represents each file parsed with information about its contents.

“We've been asked to create a package that converts a given folder of blog posts *into a* collection of Post which represents each file parsed with information about its contents.”

### [Iterative, test-driven development](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#iterative-test-driven-development)

> Iterative means we work in "thin" vertical slices

I’d remove the scare quotes around thin. Also, I don’t know what makes these slices vertical. Maybe there’s a metaphor or background I’m not familiar with. (Yup, I googled, and [there is](https://en.wikipedia.org/wiki/Vertical_slice). Maybe link or explain this for noobs like me?

> We should not trust our over-active imaginations when we start work. *You* could be

“We should not trust our over-active imaginations when we start work. *We* could be”

### [Write the test first]()

> Notice that the package of our test is blogposts_test. Remember when TDD is practiced well we take a consumer-driven approach, we don't want to test internal details because consumers don't care about that. By appending _test to our intended package name it means we can only access exported members from our package, just like a real user of our package.

I’ve seen this naming convention in other Go code, but it’s not how you do things in the earlier chapters I’ve read. That’s fine, but I wanted to flag it in case you want someday to make everything consistent. (Also, depending on where this goes in the current list, you may want to add a sentence here saying that you’ve switched things up here?)

### [Write the minimal amount of code for the test to run and check the failing test output
](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#write-the-minimal-amount-of-code-for-the-test-to-run-and-check-the-failing-test-output)

> I was surprised by the order of imports. The canonical style would put the local import last and after a skipped line, right? (This is what `goimports` does, at least.)

### [Write enough code to make it pass](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#write-enough-code-to-make-it-pass)

I was confused when you mentioned Denise. I tend to read the article at hand before following links. I suppose you’re assuming that most people will follow the link firest. Maybe introduce Denise with the link? E.g., “As Denise Yu explains, we could slime the…” On the other hand, I figured it out pretty quickly.

> remember now our focus is *to make* the test pass, not changing design, so we'll ignore the error for now.

“remember now our focus is *making* the test pass, not changing design, so we'll ignore the error for now.”

### [Error handling](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#error-handling)

> To do this "properly", we'd need a new test where we inject a failing fs.FS to make fs.ReadDir return an error.

I don’t practice TDD normally, and I had a difficult time recently figuring out how to force errors in order to test things. I mention this only to say that you might consider talking more about this sometime, even if not in this post.

### [Write enough code to make it pass](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#write-enough-code-to-make-it-pass-3)

Again, I don’t do TDD, and right around here, I grew frustrated because the placement and handling of the title, description, tags (and eventually body) fields kept changing back and forth as you refactored. In retrospect, that’s a good thing for me to learn: you may switch things back and forth a bunch as you refactor. My only suggestion is that you should maybe flag this explicitly and talk about it. E.g., “Notice that we keep changing how we assign the fields to the `Post` struct. That happens because…”

## [Iterating further](https://github.com/quii/learn-go-with-tests/blob/fstest/reading-files.md#iterating-further)

I suggest that you also mention here that `testing/fstest` is relatively limited. If people need a lot of filesystem tests, they may need [afero](https://github.com/spf13/afero) or real files. (I did pretty quickly. For example, my little project appends to a file, and I could not see a way to do that with `testing/fstest`.)
