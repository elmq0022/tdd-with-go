package arraysandslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("adds a variable amount of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		want := 36
		got := Sum(numbers)
		if got != want {
			t.Errorf("got %d does not equal %d want given numbers %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	n1 := []int{1, 2, 3}
	n2 := []int{4, 5, 6}
	got := SumAll(n1, n2)
	want := []int{6, 15}
	if !slices.Equal(got, want) {
		t.Errorf("got %d not equal to %d with n1: %v and n2: %v", got, want, n1, n2)
	}
}

func TestSumAllTails(t *testing.T) {
	n1 := []int{2, 3, 4}
	n2 := []int{4, 5, 6}
	want := []int{7, 11}
	got := SumAllTails(n1, n2)
	if !slices.Equal(want, got) {
		t.Errorf("want %d not equal %d for n1: %v, n2: %v", want, got, n1, n2)
	}
}
