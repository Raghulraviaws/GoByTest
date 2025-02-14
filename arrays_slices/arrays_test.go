package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := Sum(numbers)
	want := 36

	if got != want {
		t.Errorf("got %d and wanted %d and supplied numbers %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v and wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v and wanted %v", got, want)
		}
	}
	t.Run("Testing with slices with values", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 1})
		want := []int{2, 10}
		checkSums(t, got, want)

	})
	t.Run("Testing with slices that are empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}
		checkSums(t, got, want)
	})

}
