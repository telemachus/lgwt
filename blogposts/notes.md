# Notes on “Reading Files”

+ “We've been asked to create a package that converts a given folder of blog posts *and return* a collection of Post which represents each file parsed with information about its contents.” -> “We've been asked to create a package that converts a given folder of blog posts *into a* collection of Post which represents each file parsed with information about its contents.”
+ “We should not trust our over-active imaginations when we start work. *You* could be” -> “We should not trust our over-active imaginations when we start work. *We* could be”
+ “We're also at risk of our code coupling itself to a specific file-system implementation that we *don't really need to do*.” -> “We're also at risk of our code coupling itself to a specific file-system implementation that we *don't need*.”
+ Who’s Denise? (Problem with link-essay connection.)
+ “remember now our focus is *to make* the test pass, not changing design, so we'll ignore the error for now.” -> “remember now our focus is *making* the test pass, not changing design, so we'll ignore the error for now.”
+ Layout of imports?
+ Mismatch of names in “minimal amount of code” version? (`New` versus `NewPostsFromFS`)
+ Need to import `io/fs` as well as `testing/fstest`
+ No failing test in a case where we want an error but don’t get one. As someone who is *not* a TDD practitioner, I don’t know how to *force* an error.
