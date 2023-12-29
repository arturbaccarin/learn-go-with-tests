package arraysandslices

import (
	"reflect"
	"testing"
)

// Remember that we must not neglect our test code in the refactoring stage
func TestSum(t *testing.T) {
	t.Run("Add first 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Add other 5 numbers", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}

		got := Sum(numbers)
		want := 150

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

/* It is important to question the value of your tests.
It should not be a goal to have as many tests as possible, but rather
to have as much confidence as possible in your code base.
Having too many tests can turn in to a real problem and it just adds more overhead in maintenance.
Every test has a cost.

In our case, you can see that having two tests for this function is redundant.
If it works for a slice of one size it's very likely it'll work for a slice of any size (within reason).
*/

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Go does not let you use equality operators with slices: got != want
	// DeepEqual: verify types and values
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// The tail of a collection is all items in the collection except the first one (the "head").

// Compile time errors are our friend because they help us write software that works.
// Runtime errors are our enemies because they affect our users.
func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
