package arrayslice

import (
	"testing"
)

func sameSliceInts(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}

	return true
}

func TestSum(t *testing.T) {
	t.Run("an array of 2 ints", func(t *testing.T) {
		ns := [2]int{1, 3}
		actual := Sum(ns[:]...)
		expected := 4
		if actual != expected {
			t.Errorf("expected %d; actual %d; given %v", expected, actual, ns)
		}
	})
	t.Run("slice of five ints", func(t *testing.T) {
		ns := []int{1, 2, 3, 4, 5}
		actual := Sum(ns...)
		expected := 15

		if actual != expected {
			t.Errorf("expected %d; actual %d; given %v", expected, actual, ns)
		}
	})

	t.Run("an empty slice", func(t *testing.T) {
		ns := []int{}
		actual := Sum(ns...)
		expected := 0

		if actual != expected {
			t.Errorf("expected %d; actual %d; given %v", expected, actual, ns)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("two slices of two ints", func(t *testing.T) {
		expected := []int{3, 9}
		actual := SumAll([]int{1, 2}, []int{0, 9})

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("two empty slices", func(t *testing.T) {
		expected := []int{0, 0}
		actual := SumAll([]int{}, []int{0})

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("one empty slice", func(t *testing.T) {
		expected := []int{0}
		actual := SumAll([]int{})

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("nothing", func(t *testing.T) {
		expected := []int{}
		actual := SumAll()

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("two slices of three ints", func(t *testing.T) {
		expected := []int{3, 9}
		actual := SumAllTails([]int{10, 1, 2}, []int{10, 0, 9})

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("three slices of one int", func(t *testing.T) {
		expected := []int{0, 0, 0}
		actual := SumAllTails([]int{10}, []int{1}, []int{200})

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("four slices of various sizes", func(t *testing.T) {
		expected := []int{3, 0, 200, 14}
		actual := SumAllTails([]int{10, 2, 1}, []int{1}, []int{1, 200}, []int{1,2,3,4,5})

		if !sameSliceInts(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})
}
