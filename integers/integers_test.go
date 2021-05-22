package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 4
	actual := Add(2, 2)

	if actual != expected {
		t.Errorf("expected: '%d'; actual: '%d'", expected, actual)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
