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

func TestSumAllTails(t *testing.T) {

	assertResult := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertResult(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertResult(t, got, want)
	})

	t.Run("safely sum nil slices", func(t *testing.T) {
		got := SumAllTails(nil, []int{3, 4, 5})
		want := []int{0, 9}
		assertResult(t, got, want)
	})
}

func BenchmarkSumAllTails(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SumAllTails([]int{1, 2, 3, 4, 5}, []int{0, 9, 1, 2, 3, 4, 5, 6, 6, 7, 8})
	}
}
