package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeatTests := []struct {
		name     string
		s        string
		n        int
		expected string
	}{
		{"repeat 'a' 5 times", "a", 5, "aaaaa"},
		{"repeat 'a' 0 times", "a", 0, ""},
		{"repeat '' 5 times", "", 5, ""},
		{"repeat 'foo', 3 times", "foo", 3, "foofoofoo"},
	}


	for _, tt := range repeatTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Repeat(tt.s, tt.n)
			if actual != tt.expected {
				t.Errorf("expected %q; actual %q", tt.expected, actual)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 8))
	// Output:
	// aaaaaaaa
}
