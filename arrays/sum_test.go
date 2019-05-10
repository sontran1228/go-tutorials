package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertResult := func(t *testing.T, numbers []int, got, want int) {

		t.Helper()

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	}

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertResult(t, numbers, got, want)

	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		assertResult(t, numbers, got, want)
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
