package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/telemachus/lgwt/blogposts"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World!`
	secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte(firstBody)},
		"hola-world.md":  {Data: []byte(secondBody)},
	}

	posts, err := blogposts.New(fs)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("create a post for each file in the filesystem", func(t *testing.T) {
		expected := len(fs)
		actual := len(posts)

		if expected != actual {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("parses the title correctly", func(t *testing.T) {
		expected := "Post 1"
		actual := posts[0].Title

		if expected != actual {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})

	t.Run("parses the description correctly", func(t *testing.T) {
		expected := "Description 1"
		actual := posts[0].Description

		if expected != actual {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})

	t.Run("extracts tags from a post", func(t *testing.T) {
		expected := []string{"tdd", "go"}
		actual := posts[0].Tags

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})

	t.Run("extracts the body", func(t *testing.T) {
		expected := `Hello
World!`
		actual := posts[0].Body

		if expected != actual {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})
}
