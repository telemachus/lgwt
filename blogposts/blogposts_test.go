package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/telemachus/lgwt/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte("Title: Hello")},
		"hola-world.md":  {Data: []byte("Title: Hola")},
	}

	posts, err := blogposts.New(fs)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("create a post for each file in the filesystem", func(t *testing.T) {
		actual := len(posts)
		expected := len(fs)

		if actual != expected {
			t.Errorf("actual %d; expected %d", actual, expected)
		}
	})

	t.Run("parses the title correctly", func(t *testing.T) {
		actual := posts[0].Title
		expected := "Hello"

		if actual != expected {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})
}
