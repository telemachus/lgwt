package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
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

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("two empty slices", func(t *testing.T) {
		expected := []int{0, 0}
		actual := SumAll([]int{}, []int{0})

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("one empty slice", func(t *testing.T) {
		expected := []int{0}
		actual := SumAll([]int{})

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})

	t.Run("nothing", func(t *testing.T) {
		expected := []int{}
		actual := SumAll()

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d; actual %d", expected, actual)
		}
	})
}
